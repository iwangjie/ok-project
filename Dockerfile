FROM golang:1.15 as build
WORKDIR /release
ADD . .
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -ldflags="-s -w" -o /release/app


FROM scratch as prod
#WORKDIR /app
COPY --from=build /release/app /
ENTRYPOINT ["/app"]


#ENTRYPOINT ["pwd"]

