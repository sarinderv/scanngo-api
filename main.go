package main

import (
	"fmt"
	"os"

	"github.com/scanngo-api/cmd"
)

// @title           ScanNgo API
// @version         1.0

// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

func main() {
	if err := cmd.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error from commands.Run(): %s\n", err)
		os.Exit(1)
	}
}
