package main

import (
	"ListFiles/config"
	"ListFiles/manager"
	"ListFiles/util"
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

	//主过程
	files, err := manager.ListFiles(conf.RootFolder,true)
	err = util.Write(conf.OutputFile, files)
	if err != nil {
		panic(err)
	}
	log.Printf("task finish! time cost:%v\n", time.Since(timeStart))
}