package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/maatko/gowindcss/internal/tailwind"
)

func main() {
	binaryPath := tailwind.GetBinaryPath()
	_, err := os.Stat(binaryPath)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("[!] Failed to locate `tailwindcss` binary, downloading...")
		err := tailwind.DownloadBinary()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	command := exec.Command(binaryPath, os.Args[1:]...)

	command.Dir = dir
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	command.Run()
}
