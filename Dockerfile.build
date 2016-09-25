FROM golang:1.6
#-onbuild
ADD . /go/src/github.com/dowdeswells/testapi
WORKDIR /go/src/github.com/dowdeswells/testapi
RUN go get ./...
RUN go install -x

#ENV PORT 3100
#EXPOSE 3100
