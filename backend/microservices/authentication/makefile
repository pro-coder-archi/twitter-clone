protoc-generate:
	protoc --go_out=. --go-grpc_out=. --experimental_allow_proto3_optional \
		--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative \
		./proto/authentication.proto

serve:
	nodemon --watch "./**/*.go" --signal SIGTERM --exec "go" run main.go