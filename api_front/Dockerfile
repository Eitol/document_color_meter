FROM ubuntu:19.10
RUN apt-get update
RUN apt-get install -yq  libreoffice libreofficekit-dev golang-1.13 golang-1.13-go
ADD . /go/src/app
WORKDIR /go/src/app
RUN /usr/lib/go-1.13/bin/go build -o /go/bin/app
EXPOSE 8080
ENV GIN_MODE release
CMD ["/go/bin/app"]