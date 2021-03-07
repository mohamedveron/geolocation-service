FROM golang:1.14-alpine

RUN apk add --no-cache git
RUN apk add --update make
RUN mkdir -p /app

# Move to working directory /build
WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
COPY Makefile .
RUN go mod download

# Copy the code into the container
COPY . .

RUN make build
RUN make generate

COPY data_dump.csv .

# Export necessary port
EXPOSE 9090
# Command to run when starting the container
CMD ["bin/app"]