package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "run", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "./src"
	fmt.Println("matrix game")
	if err := cmd.Run(); err != nil {
		fmt.Println("Erreur :", err)
	}
}
