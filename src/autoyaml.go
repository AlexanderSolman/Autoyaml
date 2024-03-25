package autoyaml

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/AlexanderSolman/Autoyaml/templates"
	"sigs.k8s.io/yaml"
)

func ValidFileName(filePath string) error {
    pattern := "^[a-zA-Z0-9_-]"
    regex := regexp.MustCompile(pattern)
    substrings := strings.Split(filePath, "/")
    fileName := substrings[len(substrings)-1]
    subFileName := strings.Split(fileName, ".")
    if !regex.MatchString(subFileName[0]) { return fmt.Errorf("Invalid file name. Can only contain %s", pattern) }
    if len(subFileName) != 2 || subFileName[1] != "yaml" { return fmt.Errorf("Invalid file extension: %s. Only .yaml is valid", subFileName[1]) }
    return nil
}

func CreateFile(flag *string, filePath string, command string) error {
    var (
        data []byte
        err error
    )
    fmt.Println("From createfile:", *flag) // DEBUG
    switch command {
    case "deployment":
        data, err = yaml.Marshal(autoyaml.GenerateDeployment(flag))
        if err != nil { return err }
    case "serviceaccount":
        data, err = yaml.Marshal(autoyaml.GenerateServiceAccount(flag))
        if err != nil { return err }
    case "role":
        data, err = yaml.Marshal(autoyaml.GenerateRole(flag))
        if err != nil { return err }
    case "rolebinding":
        data, err = yaml.Marshal(autoyaml.GenerateRoleBinding(flag))
        if err != nil { return err }
    case "clusterrole":
        data, err = yaml.Marshal(autoyaml.GenerateClusterRole(flag))
        if err != nil { return err }
    case "clusterrolebinding":
        data, err = yaml.Marshal(autoyaml.GenerateClusterRoleBinding(flag))
        if err != nil { return err }
    case "pod":
        data, err = yaml.Marshal(autoyaml.GeneratePod(flag))
        if err != nil { return err }
    default:
        return fmt.Errorf("Command does not match a valid Kubernetes resource")
    }
    
    f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil { return fmt.Errorf("Failed to open file: %v", err) }
    defer f.Close()

    writer := bufio.NewWriter(f)
    err = os.WriteFile(filePath, data, 0644)
    if err != nil { return fmt.Errorf("Failed to write to file: %v", err) }
    
    err = writer.Flush()
    if err != nil { return fmt.Errorf("Failed to flush writer: %v", err) }

    return nil
}

