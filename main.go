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
	"github.com/google/credstore/client"
	"github.com/google/go-microservice-helpers/client"
	"github.com/google/go-microservice-helpers/server"
	"github.com/google/go-microservice-helpers/tracing"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	opentracing "github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/google/lvmd/proto"
	"github.com/google/lvmd/server"
)

var (
	credStoreAddress = flag.String("credstore-address", "", "credstore grpc address")
	credStoreCA      = flag.String("credstore-ca", "", "credstore server ca")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	err := tracing.InitTracer(*serverhelpers.ListenAddress, "lvmd")
	if err != nil {
		glog.Fatalf("failed to init tracing interface: %v", err)
	}

	svr := server.NewServer()
	var grpcServer *grpc.Server

	if *credStoreAddress != "" {
		appTok, err := client.GetAppToken()
		if err != nil {
			glog.Fatalf("failed to get app token: %v", err)
		}
		conn, err := clienthelpers.NewGRPCConn(*credStoreAddress, *credStoreCA, "", "")
		if err != nil {
			glog.Fatalf("failed to create connection to credstore: %v", err)
		}
		credStoreKey, err := client.GetSigningKey(context.Background(), conn, appTok)
		if err != nil {
			glog.Fatalf("failed to get signing key: %v", err)
		}

		glog.Infof("enabled credstore auth")
		grpcServer = grpc.NewServer(
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer()),
				grpc_prometheus.UnaryServerInterceptor,
				client.CredStoreTokenInterceptor(credStoreKey),
				client.CredStoreMethodAuthInterceptor(),
			)))
	} else {
		grpcServer = grpc.NewServer(
			grpc.UnaryInterceptor(
				otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer())))
	}

	pb.RegisterLVMServer(grpcServer, &svr)
	reflection.Register(grpcServer)
	grpc_prometheus.Register(grpcServer)

	err = serverhelpers.ListenAndServe(grpcServer, nil)
	if err != nil {
		glog.Fatalf("failed to serve: %v", err)
	}
}
