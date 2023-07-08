FROM golang:1.19
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go mod download
RUN go build
CMD [ "/app/github_cli" ]
