package main

import (
	"fmt"
	"runtime"

	"github.com/jaedsadadotme/golang-workshop/lesson5-export/export"
	// "export"
)

func main() {
	export.DemoExport()
	// export.demoExport() // this error because is unexport

	fmt.Println(runtime.NumCPU() + 1) // count of core cpus
}
