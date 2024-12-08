package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3deployment"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkStackProps struct {
	awscdk.StackProps
}

func NewCdkStack(scope constructs.Construct, id string, props *CdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Deploy backend lambda function and API
	lambdaFunction := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("merge-backend-function"), &awscdklambdagoalpha.GoFunctionProps{
		Entry:   jsii.String("../backend"),
		Runtime: awslambda.Runtime_PROVIDED_AL2(),
	})

	hostedZone := awsroute53.HostedZone_FromHostedZoneAttributes(stack, jsii.String("personal-website-hosted-zone"), &awsroute53.HostedZoneAttributes{
		ZoneName:     jsii.String("ofejiro.com"),
		HostedZoneId: jsii.String("Z05428331P8WLQ7347J2U"),
	})

	usEastDomainCertificate := awscertificatemanager.Certificate_FromCertificateArn(stack, jsii.String("personal-website-ssl-certificate"), jsii.String("arn:aws:acm:us-east-1:992382640465:certificate/e59b07a3-4fe3-4170-a6a9-8332cf2811ac"))
	caCentralDomainCertificate := awscertificatemanager.Certificate_FromCertificateArn(stack, jsii.String("personal-website-ca-central-ssl-certificate"), jsii.String("arn:aws:acm:ca-central-1:992382640465:certificate/8789bcda-811d-493b-a56f-65f06e8d6e6e"))

	api := awsapigateway.NewLambdaRestApi(stack, jsii.String("merge-backend-api-gateway"), &awsapigateway.LambdaRestApiProps{
		Handler: lambdaFunction,
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowOrigins: awsapigateway.Cors_ALL_ORIGINS(),
			AllowMethods: awsapigateway.Cors_ALL_METHODS(),
		},
		DomainName: &awsapigateway.DomainNameOptions{
			DomainName:  jsii.String("merge-back." + *hostedZone.ZoneName()),
			Certificate: caCentralDomainCertificate,
		},
	})

	awsroute53.NewARecord(stack, jsii.String("merge-backend-dns-a-record"), &awsroute53.ARecordProps{
		Zone:       hostedZone,
		Target:     awsroute53.RecordTarget_FromAlias(awsroute53targets.NewApiGateway(api)),
		RecordName: jsii.String("merge-back"),
	})

	app := api.Root()

	app.AddMethod(jsii.String("GET"), nil, nil)
	app.AddMethod(jsii.String("POST"), nil, nil)

	awscdk.NewCfnOutput(stack, jsii.String("merge-backend-api-gateway-endpoint"), &awscdk.CfnOutputProps{
		ExportName: jsii.String("merge-backend-api-gateway-endpoint-url"),
		Value:      api.Url(),
	})

	// Deploy frontend as a stack website on an s3 bucket + cloudfront
	staticBucket := awss3.NewBucket(stack, jsii.String("merge-website-static-bucket-named"), &awss3.BucketProps{
		AccessControl: awss3.BucketAccessControl_PRIVATE,
		// BucketName:    jsii.String("personal-website-static"),
	})

	awss3deployment.NewBucketDeployment(stack, jsii.String("merge-website-static-bucket-deployment"), &awss3deployment.BucketDeploymentProps{
		DestinationBucket: staticBucket,
		Sources:           &[]awss3deployment.ISource{awss3deployment.Source_Asset(jsii.String("../frontend/out"), &awss3assets.AssetOptions{})},
	})

	originAccessIdentity := awscloudfront.NewOriginAccessIdentity(stack, jsii.String("merge-website-origin-access-identity"), &awscloudfront.OriginAccessIdentityProps{})

	staticBucket.GrantRead(originAccessIdentity, nil)

	cloudfrontDeployment := awscloudfront.NewDistribution(stack, jsii.String("merge-website-cloudfront-deployment"), &awscloudfront.DistributionProps{
		DefaultRootObject: jsii.String("index.html"),
		DomainNames:       &[]*string{jsii.String("merge." + *hostedZone.ZoneName())},
		Certificate:       usEastDomainCertificate,
		DefaultBehavior: &awscloudfront.BehaviorOptions{
			Origin: awscloudfrontorigins.S3BucketOrigin_WithOriginAccessIdentity(staticBucket, &awscloudfrontorigins.S3BucketOriginWithOAIProps{
				OriginAccessIdentity: originAccessIdentity,
			}),
		},
	})

	awsroute53.NewARecord(stack, jsii.String("merge-website-dns-a-record"), &awsroute53.ARecordProps{
		Zone:       hostedZone,
		RecordName: jsii.String("merge"),
		Target:     awsroute53.RecordTarget_FromAlias(awsroute53targets.NewCloudFrontTarget(cloudfrontDeployment)),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkStack(app, "threewaymergecdk", &CdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
