package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

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

	fmt.Println(os.Args[1:])
	command := exec.Command(binaryPath, os.Args[1:]...)

	command.Dir = dir
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err = command.Start()
	if err != nil {
		fmt.Println("Failed to start tailwindcss command:", err.Error())
		return
	}

	fmt.Println("TailwindCSS process started. Press Ctrl+C to stop.")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan // Wait for a signal

	if command.Process != nil {
		err = command.Process.Signal(syscall.SIGTERM)
		if err != nil {
			command.Process.Kill()
		}
		command.Process.Wait()
	}
}
