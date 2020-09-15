package main

import (
	"./config"
	"./manager"
	"./util"
	"log"
	"os"
	"time"
)

func main(){
	//读配置
	conf := config.Config{}
	if !util.IsFileExist(config.ConfigFile) {
		conf.Init()
		err := conf.Save()
		if err != nil {
			panic(err)
		}
		log.Printf("please update %s\n", config.ConfigFile)
		os.Exit(0)
	}
	err := conf.Load()
	if err != nil {
		panic(err)
	}

	timeStart := time.Now()
	//读过滤文件
	var filterFileMap map[string]bool
	if util.IsFileExist(conf.FilterFile) {
		filterFileMap, err = util.ReadAllLines(conf.FilterFile)
	}

	//主过程
	files, err := manager.ListFiles(conf.RootFolder, filterFileMap, true)
	err = util.Write(config.OutputFile, files)
	if err != nil {
		panic(err)
	}
	timeEnd := time.Now()
	log.Printf("task finish! time cost:%.1f s\n", timeEnd.Sub(timeStart).Seconds())
}