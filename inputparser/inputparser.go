package inputparser

import (
	"bufio"
	"log"
	"os"
)

const (
	KB = 1024
	MB = KB * 1024
)

func ParseCLI() string {
	scanner := bufio.NewScanner(os.Stdin)
	content := ""
	for scanner.Scan() {
		l := scanner.Text()
		if l == "exit" {
			break
		}
		content += "\n" + l
	}
	return content
}

func ParseFile(filename string) string {
	info, err := os.Stat(filename)
	if err != nil {
		log.Panicln("parsing error:", err)
	}
	if info.Size() > MB*5 {
		log.Panicln("File too big to parse. File size should be kept under 5 MB")
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Panicln("parsing error:", err)
	}
	return string(content)
}
