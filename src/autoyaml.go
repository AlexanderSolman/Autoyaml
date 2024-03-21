package autoyaml

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/AlexanderSolman/Autoyaml/templates"
)

func validFileName(filePath string) error {
    pattern := "^[a-zA-Z0-9_-]"
    regex := regexp.MustCompile(pattern)
    substrings := strings.Split(filePath, "/")
    fileName := substrings[len(substrings)-1]
    subFileName := strings.Split(fileName, ".")
    if !regex.MatchString(subFileName[0]) { return fmt.Errorf("Invalid file name. Can only contain %s", pattern) }
    if len(subFileName) != 2 || subFileName[1] != "yaml" { return fmt.Errorf("Invalid file extension: %s. Only .yaml is valid", subFileName[1]) }
    return nil
}

func createDeployment(flag *string, filePath string) error {
    
    // ~/project/src/deployment/nginx.yaml || nginx.yaml
    

    deployment := autoyaml.GenerateDeployment(flag)    

}

func createSeriveAccount(flag *string, filePath string) error {
    
    serviceaccount := autoyaml.GenerateServiceAccount(flag)
}

func createRole(flag *string, filePath string) error {
    
    createrole := autoyaml.GenerateRole(flag)
}

func createRoleBinding(flag *string, filePath string) error {

    rolebinding := autoyaml.GenerateRoleBinding(flag)
}

func createClusterRole(flag *string, filePath string) error {

    clusterrole := autoyaml.GenerateClusterRole(flag)
}

func createClusterRoleBinding(flag *string, filePath string) error {

    clusterrolebinding := autoyaml.GenerateClusterRoleBinding(flag)
}

func createPod(flag *string, filePath string) error {

    pod := autoyaml.GeneratePod(flag)
}

