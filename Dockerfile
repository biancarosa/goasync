FROM golang:1.14.2 AS build

WORKDIR $GOPATH/src/github.com/biancarosa/goasync
COPY . ./
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY  --from=build /app ./
ENTRYPOINT ["./app"]