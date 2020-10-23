package manager

import (
	"ListFiles/config"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//列出目录内的文件
func ListFiles(folder string, filterFileMap map[string]bool, isProjectFolder bool) ([]string, error){
	files := make([]string, 0)
	fis, err := ioutil.ReadDir(folder)
	if err != nil {
		return files, err
	}
	cnt := 0
	for _, fileInfo := range fis {
		if isProjectFolder {
			cnt++
			log.Printf("%.2f%%\n", float64(cnt*100)/float64(len(fis)))
		}

		path := folder + string(os.PathSeparator) + fileInfo.Name()
		filename := fileInfo.Name()
		if fileInfo.IsDir() {
			fs, err := ListFiles(path, filterFileMap, false)
			if err != nil {
				continue
			}
			files = append(files, fs...)
		}else{
			if !strings.Contains(filename, "."){
				continue
			}
			suffix := filename[strings.LastIndex(filename, ".")+1:]
			suffix = strings.ToLower(suffix)
			_, ok := config.FileSuffixMap[suffix]
			if ok {
				projectPath := path[len(config.RootFolder)+1:]
				_, ok2 := filterFileMap[projectPath]
				if !ok2 {
					files = append(files, path)
				}
			}
		}
	}
	return files, err
}
