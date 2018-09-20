.PHONY: prepare
prepare:
	brew install protoc

.PHONY: prod-prepare
prod-prepare:
	PROTOC_ZIP=protoc-3.3.0-linux-x86_64.zip
	curl -OL https://github.com/google/protobuf/releases/download/v3.3.0/$PROTOC_ZIP
	sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
	rm -f $PROTOC_ZIP

.PHONY: build-grpc-todo
build-grpc-todo:
	protoc -I. --go_out=plugins=grpc:. api/proto/v1/todo-service.proto

.PHONY: build-grpc-topic
build-grpc-topic:
	protoc -I. --go_out=plugins=grpc:. api/proto/v1/topic.proto	