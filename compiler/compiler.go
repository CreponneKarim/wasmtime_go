package compiler

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Compile(filePath string) (string, error){
	//	compile
	slices := strings.Split(filePath, "/")
	if len(slices)<1 {
		return "", fmt.Errorf("invalid file path")
	}
	fileName := strings.Split(slices[len(slices)-1], ".")[0]
	resultFileName := fileName + ".wasm"
	args := []string{"build", "-o", resultFileName, filePath}
	cmd := exec.Command("GOARCH=wasm GOOS=wasip1 go",args...)
	//	pass it to execution

	output, err := cmd.CombinedOutput()

    if err != nil {
    	return "", fmt.Errorf("Failed to compile Go code: %v\nOutput:\n%s", err, output)
    }

    log.Printf("Successfully compiled Go code.\nOutput:\n%s", output)

	return resultFileName, nil
}