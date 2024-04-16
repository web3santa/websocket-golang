FROM golang:1.21.5 AS build

WORKDIR /app

COPY go.mod .
COPY main.go .
COPY hub.go .
COPY client.go .
COPY templates ./templates

RUN go get
RUN go build -o server .

FROM nginx:latest

WORKDIR /usr/share/nginx/html

COPY --from=build /app/templates/ /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]
