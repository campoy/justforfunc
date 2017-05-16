FROM alpine

RUN apk update && apk add flite
ADD app /app
ENTRYPOINT ["/app"]