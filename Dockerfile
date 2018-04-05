FROM golang:latest AS build

WORKDIR $GOPATH/src/github.com/biancarosa/goasync
COPY . ./
RUN go get github.com/tools/godep
RUN godep get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY  --from=build /app ./
ENTRYPOINT ["./app"]