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
    var (
        err error
        flag string
        filePath string
    )
    if len(os.Args) < 3 {
        fmt.Println("Not enough arguments")
        os.Exit(1)
    } else if len(os.Args) > 4 {
        fmt.Println("Too many arguments")
        os.Exit(1)
    }

    amountOfArguments := len(os.Args)

    command := strings.ToLower(os.Args[1])
    if amountOfArguments == 4 { 
        flag = os.Args[2]
        if flag != "-a" {
            var i string
            fmt.Printf("Invalid flag: %s. Did you mean -a? [y,n]", flag)
            fmt.Scanln(&i)
            if i == "y" { 
                flag = "-a" 
            } else { 
                flag = ""
                fmt.Println("Flag is set to deafult")
            }
        }
        filePath = os.Args[3] // Should be absolute or defaults current directory (name.yaml has to be set)
    } else {
        flag = ""
        filePath = os.Args[2]
    }
    
    err = validFileName(filePath)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    err = createFile(&flag, filePath, command)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

}
