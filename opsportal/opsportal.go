package opsportal

import (
	"fmt"

	pb "github.com/HierIsPierre/nexus/opsportal/protos"
)

func DeploySandbox(subscription string, services []string) error {
	req := &pb.DeploySandboxRequest{
		SubscriptionName: subscription,
	}

	fmt.Println(
		req.SubscriptionName,
	)
	return nil
}

func RemoveSandbox(subscription string) error {

	return nil
}
