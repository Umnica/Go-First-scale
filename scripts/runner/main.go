package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"io/ioutil"
	"strings"
)

func main() {
	scriptsDir := "scripts"
	
	files, err := ioutil.ReadDir(scriptsDir)
	if err != nil {
		fmt.Println("Error reading scripts directory:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run runner/main.go <script_name>")
		fmt.Println("Available scripts:")
		for _, f := range files {
			if !f.IsDir() && strings.HasSuffix(f.Name(), ".go") {
				fmt.Println("  -", f.Name())
			}
		}
		return
	}

	scriptName := os.Args[1]
	if !strings.HasSuffix(scriptName, ".go") {
		scriptName += ".go"
	}

	scriptPath := filepath.Join(scriptsDir, scriptName)
	
	cmd := exec.Command("go", "run", scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running script: %v\n", err)
		os.Exit(1)
	}
}
