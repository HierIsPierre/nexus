package docker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Build(service string, name string) {
	dockerFilePath := fmt.Sprintf("%v/src/%v.Presentation/Dockerfile", service, service)
	imageTag := fmt.Sprintf("psiclesaasservicesregistry.azurecr.io/%v:%s", strings.ToLower(service), name)

	fmt.Printf("Image Tag: %v\n\n", imageTag)
	cmd := exec.Command("docker", "build", "-f", dockerFilePath, ".", "-t", imageTag) 
	fmt.Printf("Running command: %v", cmd.Args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error building image: %v\n", err)
	}

	Push(imageTag)
}

func Push(imageTag string) {
	cmd := exec.Command("docker", "push", imageTag)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error push image: %v\n", err)
	}
}

