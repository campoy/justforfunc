FROM alpine

RUN apk update && apk add flite
ADD app /
ENTRYPOINT ["/app"]