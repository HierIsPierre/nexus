package opsportal

import (
	"fmt"

	pb "github.com/HierIsPierre/opsportal/protos"
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
