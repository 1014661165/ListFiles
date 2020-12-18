package config

import (
	"encoding/xml"
	"io/ioutil"
	"strings"
)

const (
	ConfigFile = "./list-files.xml"
)
var (
	FileSuffixMap = make(map[string]bool)
	RootFolder string
)

type Config struct {
	RootFolder string `xml:"root_folder"`
	FileSuffix FileSuffix `xml:"file_suffix"`
	OutputFile string `xml:"output_file"`
}

type FileSuffix struct {
	Suffix []string `xml:"suffix"`
}

func (c *Config) Init(){
	c.RootFolder = ""
	c.FileSuffix = FileSuffix{Suffix:[]string{""}}
}

func (c *Config) Save() error {
	content, _ := xml.MarshalIndent(c, "", "\t")
	err := ioutil.WriteFile(ConfigFile, content, 0744)
	return err
}

func (c *Config) Load() error{
	content, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(content, c)
	RootFolder = c.RootFolder
	if err == nil {
		for _,s := range c.FileSuffix.Suffix {
			suffix := strings.ToLower(s)
			FileSuffixMap[suffix] = true
		}
	}
	return err
}