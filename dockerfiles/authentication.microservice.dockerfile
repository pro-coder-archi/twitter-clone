FROM go:1.19-alpine AS builder

WORKDIR /app

# resolving dependencies of modules

COPY ./backend/microservices/authentication/go.mod ./backend/microservices/authentication
COPY ./backend/microservices/authentication/go.sum ./backend/microservices/authentication

RUN cd ./backend/microservices/authentication && \
    go mod tidy

COPY ./backend/shared/go.mod ./backend/shared
COPY ./backend/shared/go.sum ./backend/shared

RUN cd ./backend/shared && \
    go mod tidy

#* copying the main module
COPY ./backend/microservices/authentication/* ./backend/microservices/authentication

#* copying the library module containing shared code
COPY ./backend/shared/* ./backend/shared

#* creating a workspace file and registering the modules
RUN go work init
RUN go work use ./backend

RUN go build -o build ./backend/microservices/authentication

FROM alpine:latest AS packager

WORKDIR /

COPY --from=builder /app/build .

ENTRYPOINT [ "build" ]