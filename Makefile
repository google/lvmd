all: grpc

grpc: proto/lvm.pb.go

proto/lvm.pb.go: proto/lvm.proto
	cd proto && protoc -I/usr/local/include -I. \
	  --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
	  lvm.proto

clean:
	rm -f proto/lvm.pb.go

.PHONY: all clean
