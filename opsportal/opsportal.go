package opsportal

import (
	"context"
	"fmt"
	"crypto/x509"
	"os"

	pb "github.com/HierIsPierre/nexus/opsportal/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func DeploySandbox(subscription string, services []string, imageName string) error {
	cert, err := os.ReadFile("../cert/aspnetcore-dev-cert.pem")
    if err != nil {
        return err
    }

    // Create a certificate pool and add your certificate
    certPool := x509.NewCertPool()
    if ok := certPool.AppendCertsFromPEM(cert); !ok {
		return err
    }

    // Create TLS credentials using the cert pool
    creds := credentials.NewClientTLSFromCert(certPool, "")
	if err != nil {
		return fmt.Errorf("Could not create creds")
	}

	conn, err := grpc.NewClient(
		"localhost:5001", grpc.WithTransportCredentials(creds),
	)
	
	if err != nil {
		return fmt.Errorf("Coulld not create a grpc connection")
	}
	
	if len(services) == 0 {
		return fmt.Errorf(
			"Please specify the services that you want to deploy to the sandbox",
		)
	}
	
	serviceDetails := make([]*pb.ServiceDetails, 0, len(services))
	for _, service := range services {
		serviceDetails = append(serviceDetails, &pb.ServiceDetails{
			ServiceName: service,
			Image: imageName,
		})	
	}

	req := &pb.DeploySandboxRequest{
		SubscriptionName: subscription,
		ServiceDetails: serviceDetails,
	}


	client := pb.NewDeploySandboxClient(conn)
	response, err := client.DeploySandbox(context.Background(), req)

	if err != nil {
		fmt.Printf("An error occurred: %s", err)
		return err
	}

	fmt.Printf("Sandbox deployed for %s %s\n", subscription, response)
	return nil
}

func RemoveSandbox(subscription string) error {

	return nil
}
