FROM golang:alpine3.19 as go_build

WORKDIR /main

COPY ./backend/go.mod ./backend/go.sum ./

RUN go mod download && go mod verify

COPY ./backend .

RUN go build

FROM node:alpine3.19 as node_package_build

WORKDIR /main

COPY ./frontend/package.json .

RUN npm i

FROM node:alpine3.19 as static_build
WORKDIR /main

COPY --from=node_package_build /node_modules ./node_modules
COPY ./frontend .

RUN npm run build

FROM busybox:1.36 as main
COPY --from=static_build /out ./website
COPY --from=go_build /three_way_merge .

EXPOSE 3000 8000

CMD [ "sh", "entrypoint.sh" ]