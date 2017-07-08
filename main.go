/*

Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/google/go-microservice-helpers/server"
	"github.com/google/go-microservice-helpers/tracing"

	pb "github.com/google/lvmd/proto"
	"github.com/google/lvmd/server"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	err := tracing.InitTracer(*serverhelpers.ListenAddress, "lvmd")
	if err != nil {
		glog.Fatalf("failed to init tracing interface: %v", err)
	}

	svr := server.NewServer()

	grpcServer, _, err := serverhelpers.NewServer()
	if err != nil {
		glog.Fatalf("failed to init GRPC server: %v", err)
	}

	pb.RegisterLVMServer(grpcServer, &svr)

	err = serverhelpers.ListenAndServe(grpcServer, nil)
	if err != nil {
		glog.Fatalf("failed to serve: %v", err)
	}
}
