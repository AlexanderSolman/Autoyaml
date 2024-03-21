package autoyaml

import (
	"fmt"
	"os"
	"strings"
)

// >> autoyaml <kind(deployment)> <file/path/your.yaml(defaults to current directory ./your.yaml)>
// or
// >> autoyaml <kind(deployment)> <flag(-b(basic)|-a(advanced))> <file/path/your.yaml(defaults to current directory ./your.yaml)>

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Not enough arguments")
        os.Exit(1)
    } else if len(os.Args) > 3 {
        fmt.Println("Too many arguments")
        os.Exit(1)
    }

    command := strings.ToLower(os.Args[1])
    flag := os.Args[2]
    filePath := os.Args[3] // Should be absolute or defaults current directory (name.yaml has to be set)
    err := validFileName(filePath)
    if err != nil {
       fmt.Println(err)
    }
    if flag != "-a" || flag != "" { fmt.Printf("Invalid flag: %s, defaults as blank", flag) }

    switch command {
    case "deployment":
        createDeployment(&flag, filePath)
    case "serviceaccount":
        createSeriveAccount(&flag, filePath)
    case "role":
        createRole(&flag, filePath)
    case "rolebinding":
        createRoleBinding(&flag, filePath)
    case "clusterrole":
        createClusterRole(&flag, filePath)
    case "clusterrolebinding":
        createClusterRoleBinding(&flag, filePath)
    case "pod":
        createPod(&flag, filePath)
    default:
        fmt.Println("Command does not match a valid Kubernetes resource")
        os.Exit(1)
    }


}
