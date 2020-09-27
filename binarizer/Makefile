.PHONY: gen dev_install

PROTO_PATH=./api/protos
OUT=./api/gen

gen:
	@mkdir -p ${OUT}/python
	@python3 -m grpc_tools.protoc -I${PROTO_PATH} --python_out=. --grpc_python_out=. ${PROTO_PATH}/*
	@mkdir -p ${OUT}/go
	@protoc -I=${PROTO_PATH} --go_out=plugins=grpc:${OUT}/go ${PROTO_PATH}/*


dev_install:
	@pip3 install grpcio grpcio-tools
	@go get -u github.com/golang/protobuf/protoc-gen-go
