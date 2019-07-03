
package main

import (
	"fmt"
	"time"
	"github.com/pepabo/go-netapp/netapp"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/prometheus/common/log"

)


type Filer struct {
	Name             string `yaml:"name"`
	Host             string `yaml:"host"`
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
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

func newNetappClient(filer *Filer) *netapp.Client {

	_url := "https://%s/servlets/netapp.servlets.admin.XMLrequest_filer"
	url := fmt.Sprintf(_url, filer.Host)

	version := "1.7"

	opts := &netapp.ClientOptions{
		BasicAuthUser:     filer.Username,
		BasicAuthPassword: filer.Password,
		SSLVerify:         false,
		Timeout:           30 * time.Second,
	}

	return netapp.NewClient(url, version, opts)
}
