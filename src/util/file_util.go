package util

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

//判断文件或目录是否存在
func IsFileExist(path string) bool{
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//读取问文件所有行
func ReadAllLines(path string) (map[string]bool, error){
	lines := make(map[string]bool)
	f, err := os.OpenFile(path, os.O_RDONLY, 0744)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || err == io.EOF{
			if line != "" {
				line = strings.TrimRight(line, " \r\n")
				lines[line] = true
			}
			break
		}
		line = strings.TrimRight(line, " \r\n")
		lines[line] = true
	}
	return lines, nil
}

//输出到文件
func Write(path string, lines []string) error{
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0744)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	for _, line := range lines {
		_, err = writer.WriteString(line + "\n")
	}
	err = writer.Flush()
	return err
}
