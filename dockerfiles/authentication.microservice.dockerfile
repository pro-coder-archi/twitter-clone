FROM go:1.19-alpine AS builder

WORKDIR /app

#* copying the main and dependent go modules
COPY ./backend/microservices/authentication/* ./backend/microservices/authentication

#* creating a workspace file and registering the modules
RUN go work init
RUN go work use ./backend

RUN go build -o build ./backend/microservices/authentication

FROM alpine:latest AS packager

WORKDIR /

COPY --from=builder /app/build .

ENTRYPOINT [ "build" ]