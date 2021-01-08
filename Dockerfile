FROM golang:1.13.5-alpine3.10 as build
RUN adduser -u 10001 -D app-runner
WORKDIR /release
ADD . .
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV GOOS=linux
RUN go build -o app


FROM alpine:3.10 as prod
#COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /release/app /
#USER app-runner
ENTRYPOINT ["/app"]


