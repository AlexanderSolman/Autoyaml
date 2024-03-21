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
    var err error
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
    err = validFileName(filePath)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    if flag != "-a" || flag != "" { fmt.Printf("Invalid flag: %s, defaults as blank", flag) }

    err = createFile(&flag, filePath, command)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

}
