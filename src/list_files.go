package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	_ "io/ioutil"
	"os"
	"strings"
	"time"
)

func Exist(path string) bool{
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func IsDir(path string) bool{
	fileInfo, _ := os.Stat(path)
	return fileInfo.IsDir()
}

func Contains(exts []string, target string) bool{
	contains := false
	for _, ext := range exts{
		if ext == target{
			contains = true
			break
		}
	}
	return contains
}

func GetAllFiles(dir string, list *list.List) error {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil{
		return err
	}

	for _, file := range fileInfos{
		path := dir + string(os.PathSeparator) + file.Name()
		if file.IsDir(){
			_ = GetAllFiles(path, list)
		}else{
			if len(os.Args) > 2{
				ext := file.Name()[strings.LastIndex(file.Name(), ".")+1:]
				contains := Contains(os.Args[2:], ext)
				if contains{
					list.PushBack(path)
				}
			}else{
				list.PushBack(path)
			}
		}
	}
	return nil
}

func CheckArgs(){
	if len(os.Args) < 2{
		fmt.Println("usage: ./*.exe dir ext1 ext2 ...")
		os.Exit(0)
	}
	if !Exist(os.Args[1]){
		fmt.Printf("%s not exist", os.Args[1])
		os.Exit(0)
	}

	if !IsDir(os.Args[1]){
		fmt.Printf("%s is not a directory", os.Args[1])
		os.Exit(0)
	}
}

func WriteFileList(fileList *list.List){
	outputFile, err := os.OpenFile("files.txt", os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil{
		panic(err)
	}
	defer outputFile.Close()
	for path := fileList.Front(); path != nil; path = path.Next(){
		outputFile.WriteString(path.Value.(string) + "\n")
	}
}


func main() {
	CheckArgs()
	fileList := list.New()
	fmt.Println("loading files...")
	start := time.Now()
	err := GetAllFiles(os.Args[1], fileList)
	if err != nil{
		panic(err)
	}
	fmt.Println("saving...")
	WriteFileList(fileList)
	end := time.Now()
	fmt.Printf("task finish! time cost %.2fs\n", end.Sub(start).Seconds())
}
