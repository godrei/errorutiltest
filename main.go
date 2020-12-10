package main

import (
	"fmt"

	"github.com/bitrise-io/go-utils/errorutil"
)

func main() {
	i, _ := errorutil.CmdExitCodeFromError(nil)
	fmt.Println(i)
}
