FROM golang:1.21 AS build

WORKDIR /app 

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build \
  -ldflags="-linkmode external -extldflags -static" \
  -tags netgo \
  -o web-app-golang

FROM scratch

WORKDIR /

COPY --from=build /app/web-app-golang web-app-golang

CMD ["/web-app-golang"]