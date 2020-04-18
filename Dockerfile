FROM golang:alpine
ARG PKGNAME
RUN apk --no-cache --update upgrade && apk --no-cache add git gcc libc-dev ca-certificates
ENV GO111MODULE=on
ENV PKGNAME=${PKGNAME}
RUN mkdir -p /opt/code
WORKDIR /opt/code
ENTRYPOINT [ "./test.sh" ]
