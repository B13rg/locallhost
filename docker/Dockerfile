FROM golang:1.24.2 as builder

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ./locallhost

FROM gcr.io/distroless/static-debian12 AS result

COPY --from=builder /app/locallhost /locallhost

USER nonroot:nonroot

EXPOSE 8080

# Run
CMD ["/locallhost", "-p", "8080"]