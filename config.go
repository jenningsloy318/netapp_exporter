package main

import (
	"fmt"
	"github.com/pepabo/go-netapp/netapp"
	"github.com/prometheus/common/log"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type Filer struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Debug    bool   `yaml:"debug"`
}

func loadFilerFromFile(fileName string) (c []*Filer) {
	var fb []Filer
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("[ERROR] ", err)
	}
	err = yaml.Unmarshal(yamlFile, &fb)
	if err != nil {
		log.Fatal("[ERROR] ", err)
	}
	for _, b := range fb {
		c = append(c, &b)
	}
	return
}

func newNetappClient(filer *Filer) (string, *netapp.Client) {

	_url := "https://%s/servlets/netapp.servlets.admin.XMLrequest_filer"
	url := fmt.Sprintf(_url, filer.Host)

	version := "1.130"

	opts := &netapp.ClientOptions{
		BasicAuthUser:     filer.Username,
		BasicAuthPassword: filer.Password,
		SSLVerify:         false,
		Debug:						 filer.Debug,
		Timeout:           30 * time.Second,
	}
	netappClient :=netapp.NewClient(url, version, opts)
	return filer.Name, netappClient
}
