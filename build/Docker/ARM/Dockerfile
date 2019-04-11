# build image
FROM golang as builder

WORKDIR /opt

COPY . .

# compile
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -o app cmd/app/* 

# deployment image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

LABEL author="Stephen Onnen"

WORKDIR /root/

COPY --from=builder /opt/ .

CMD [ "./app" ]

EXPOSE 8080