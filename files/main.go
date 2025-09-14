package main

import (
	"fmt"
	fileutils "go/files/filemanager"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/text.txt"
	c, err := fileutils.ReadFile(filePath)
	if err == nil {
		fmt.Println(c)
		newContent := fmt.Sprintf("%v \nNew line added!!", c)
		fileutils.WriteFiles(filePath+".answer.txt", newContent)
	} else {
		fmt.Printf("Ouch!! %v", err)
	}
}
