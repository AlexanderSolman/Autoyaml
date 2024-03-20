package autoyaml

import (
    "fmt"
    "os"
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



}
