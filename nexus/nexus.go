package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/HierIsPierre/nexus/docker"
	"github.com/HierIsPierre/nexus/opsportal"
)

type stringSlice []string

func (s *stringSlice) String() string {
	return fmt.Sprint(*s)
}

func (s *stringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'deploy' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "sandbox":
		deployCmd := flag.NewFlagSet("sandbox", flag.ExitOnError)
		var imageName string
		var subscriptionName string
		var services stringSlice

		deployCmd.StringVar(&imageName, "imagename", "", "Name of the image")
		deployCmd.StringVar(&subscriptionName, "subscription", "", "Name of the subscription to which to deploy")
		deployCmd.Var(&services, "service", "Service to deploy (can be repeated)")

		// Parse flags after "deploy"
		deployCmd.Parse(os.Args[2:])

		// if len(services) == 0 {
		// 	fmt.Println("Error: -service flag is required")
		// 	deployCmd.Usage()
		// }

		if len(subscriptionName) == 0 {
			fmt.Println("Error: -subscription flag is required")
			deployCmd.Usage()
			return
		}

		for _, service := range services {
			fmt.Printf("Deploying service: %s (name: %s)\n", service, imageName)
			docker.Build(service, imageName)
		}

		sandboxImage := imageName + "-sandbox"

		err := opsportal.DeploySandbox(
			subscriptionName,
			services,
			sandboxImage,
		)

		if err != nil {
			fmt.Printf("%s\n\n", err)
		}

	default:
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println("Available command: sandbox")
	}
}
