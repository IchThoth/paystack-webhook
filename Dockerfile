FROM go1.21alpine as builder

RUN go mod download

