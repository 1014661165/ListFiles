package config

import (
	"encoding/xml"
	"io/ioutil"
	"strings"
)

const (
	ConfigFile = "./ListFiles.xml"
)
var (
	FileSuffixMap = make(map[string]bool)
)

type Config struct {
	RootFolder string `xml:"root_folder"`
	FileSuffix FileSuffix `xml:"file_suffix"`
	FilterFile string `xml:"filter_file"`
	OutputFile string `xml:"output_file"`
}

type FileSuffix struct {
	Suffix []string `xml:"suffix"`
}

func (c *Config) Init(){
	c.RootFolder = ""
	c.FileSuffix = FileSuffix{Suffix:[]string{""}}
	c.FilterFile = ""
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
	if err == nil {
		for _,s := range c.FileSuffix.Suffix {
			suffix := strings.ToLower(s)
			FileSuffixMap[suffix] = true
		}
	}
	return err
}