package fileutils

import "os"

func WriteFiles(filename string, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}
