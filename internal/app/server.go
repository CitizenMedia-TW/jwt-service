package app

import (
	"flag"
	"fmt"
	"jwt-service/internal/config"
	"jwt-service/internal/grpcapp"
	"jwt-service/internal/restapp"
	"log"
	"net"
	"net/http"
)

var (
	grpcport = flag.Int("grpcport", 50052, "The grpc server port")
	restport = flag.String("restport", ":7999", "The rest server port")
)

func StartServer(cnf config.Config) {
	/*
	 * Start the gRPC server
	 */
	go func() {
		grpca := grpcapp.New(cnf)
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcport))
		if err = grpca.Server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	/*
	 * Start the REST server
	 */
	resta := restapp.New(cnf)
	srv := &http.Server{
		Addr:    *restport,
		Handler: resta.Routes(),
	}

	log.Printf("Server started successfully!\n\tRunning gRPC on port %d and REST on port %s", *grpcport, *restport)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
