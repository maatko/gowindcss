package tailwind

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func GetDownloadLink() string {
	os := runtime.GOOS
	if os == "darwin" {
		os = "macos"
	}

	arch := runtime.GOARCH
	if arch == "amd64" {
		arch = "x64"
	}

	var link = fmt.Sprintf("https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-%s-%s", os, arch)
	if os == "windows" {
		link += ".exe"
	}
	return link
}

func GetBinaryPath() string {
	binaryPath := filepath.Join(os.TempDir(), "tailwindcss")
	if runtime.GOOS == "windows" {
		binaryPath += ".exe"
	}
	return binaryPath
}

func DownloadBinary() error {
	response, err := http.Get(GetDownloadLink())
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to download `tailwindcss` binary. [%d]", response.StatusCode)
	}

	binary, err := os.Create(GetBinaryPath())
	if err != nil {
		return fmt.Errorf("Failed to create `tailwindcss` file. [%s]", err.Error())
	}
	defer binary.Close()

	_, err = io.Copy(binary, response.Body)
	if err != nil {
		return fmt.Errorf("Failed writing to `tailwindcss` file. [%s]", err.Error())
	}

	os.Chmod(GetBinaryPath(), 0755)
	return nil
}
