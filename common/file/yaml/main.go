package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name           string `yaml:"name"`
	ImportPlaybook string `yaml:"import_playbook"`
	Vars           Vars   `yaml:"vars"`
}

type Vars struct {
	HostName                string `yaml:"hostname"`
	InfrawareLeaderHostName string `yaml:"infraware_leader_hostname"`
}

func main() {

	//config := Config{Record: Record{Item: "window", Col: "blue", Size: "small"}}
	//config := []Record{{Item: "window", Col: "blue", Size: "small"}}

	config := []Config{
		{
			Name:           "install-001",
			ImportPlaybook: "../playbooks/install.yml",
			Vars: Vars{
				HostName:                "in",
				InfrawareLeaderHostName: "manager",
			},
		},
		{
			Name:           "install-002",
			ImportPlaybook: "../playbooks/install.yml",
			Vars: Vars{
				HostName:                "ex",
				InfrawareLeaderHostName: "manager",
			},
		},
	}

	data, err := yaml.Marshal(&config)

	if err != nil {
		log.Fatal(err)
	}

	err2 := ioutil.WriteFile("./config.yaml", data, 0644)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("data written")
}
