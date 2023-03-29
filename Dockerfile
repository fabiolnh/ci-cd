FROM golang:latest
RUN mkdir -p /app
USER nonroot
WORKDIR /app
COPY /app .
RUN go env -w GO111MODULE=off
RUN go build -o math
CMD ["./math"]
