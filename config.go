package main

import (
	"fmt"
	"github.com/pepabo/go-netapp/netapp"
	"github.com/prometheus/common/log"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
	"sync"
)

type Config struct {
	Credentials map[string]Credential `yaml:"credentials"`
}

type SafeConfig struct {
	sync.RWMutex
	C *Config
}

type Credential struct {
	Group     string `yaml:"group"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Debug    bool   `yaml:"debug"`
}

func (sc *SafeConfig) ReloadConfig(configFile string) error {
	var c = &Config{}

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Errorf("Error reading config file: %s", err)
		return err
	}
	if err := yaml.Unmarshal(yamlFile, c); err != nil {
		log.Errorf("Error parsing config file: %s", err)
		return err
	}

	sc.Lock()
	sc.C = c
	sc.Unlock()

	log.Infoln("Loaded config file")
	return nil
}



func (sc *SafeConfig) CredentialsForTarget(target string) (*Credential, error) {
	sc.Lock()
	defer sc.Unlock()
	if credential, ok := sc.C.Credentials[target]; ok {
		return &Credential{
			Group:     credential.Group,
			Username:     credential.Username,
			Password: credential.Password,
			Debug: credential.Debug,
		}, nil
	}
	if credential, ok := sc.C.Credentials["default"]; ok {
		return &Credential{
			Group:     credential.Group,
			Username:     credential.Username,
			Password: credential.Password,
			Debug: credential.Debug,
		}, nil
	}
	return &Credential{}, fmt.Errorf("no credentials found for target %s", target)
}


func newNetappClient(host string, credential *Credential) (string, *netapp.Client) {

	_url := "https://%s/servlets/netapp.servlets.admin.XMLrequest_filer"
	url := fmt.Sprintf(_url, host)

	version := "1.130"

	opts := &netapp.ClientOptions{
		BasicAuthUser:     credential.Username,
		BasicAuthPassword: credential.Password,
		SSLVerify:         false,
		Debug:						 credential.Debug,
		Timeout:           30 * time.Second,
	}
	netappClient :=netapp.NewClient(url, version, opts)
	return credential.Group, netappClient
}










