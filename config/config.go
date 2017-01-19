package config

import (
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type DucuemuConfig struct {

	path 	string
	cfg	*ConfigType

}

type ConfigType struct {

	DB [5]string `json:"db"`
	Host string `json:"host"`
	Hostname string `json:"hostname"`
}

func NewConfig(path string) *DucuemuConfig {
	if path == "" {
		path = "./config.json"
	}
	cfg := &DucuemuConfig{path, &ConfigType{}}
	return cfg
}

func (dc *DucuemuConfig) Read(){

	file, err := ioutil.ReadFile(dc.path) // For read access.

	if err != nil {
		log.Fatal(err)
	} else {

		jsoncfg := ConfigType{}

		if err := json.Unmarshal(file, &jsoncfg); err != nil {
			fmt.Println("Could not read " + dc.path )
		}
		dc.cfg = &jsoncfg
	}
}

func (dc *DucuemuConfig) DB() (string, string, string, string, string){
	return dc.cfg.DB[0], dc.cfg.DB[1], dc.cfg.DB[2], dc.cfg.DB[3], dc.cfg.DB[4]
}

func (dc *DucuemuConfig) Host() (string){
	return dc.cfg.Host
}

func (dc *DucuemuConfig) HostName() (string){
	return dc.cfg.Hostname
}