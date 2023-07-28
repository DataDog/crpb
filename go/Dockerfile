FROM golang:1.18-buster as builder
WORKDIR /app
COPY main.go go.mod go.sum /app/

# Install Orchestrion
RUN go install github.com/datadog/orchestrion@latest
RUN orchestrion -w ./
RUN go mod tidy

RUN go build -o app

FROM debian:stable-slim
RUN apt-get update
RUN apt-get install -y ca-certificates
COPY --from=builder /app/app /app/app
COPY --from=datadog/serverless-init:1.0.0 /datadog-init /app/datadog-init

ENV DD_SERVICE=datadog-demo-run-go
ENV DD_ENV=datadog-demo
ENV DD_VERSION=1

ENTRYPOINT ["/app/datadog-init"]
CMD ["/app/app"]
