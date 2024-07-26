SHELL:=/bin/bash

proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
  		--go-grpc_opt=paths=source_relative pb/comment/*.proto

proto_php:
	@protoc --php_out=. --php_opt=paths=source_relative --php-grpc_out=. \
  		--php-grpc_opt=paths=source_relative pb/comment/*.proto