FROM golang:1.14 as build

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/app
COPY . .

RUN dep ensure
RUN CGO_ENABLED=0 go build -a -installsuffix nocgo -o app-bin .

FROM scratch
COPY  --from=build /go/src/app/app-bin ./
ENTRYPOINT ["./app-bin"]