FROM node:alpine3.19

# Get Golang
COPY --from=golang:alpine3.19 /usr/local/go/ /usr/local/go/

ENV PATH="/usr/local/go/bin:${PATH}"


WORKDIR /main

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build

EXPOSE 3000 8000

RUN mkdir frontend/node_modules

VOLUME [ "frontend/node_modules" ]

CMD [ "sh", "entrypoint.sh" ]