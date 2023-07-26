package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/apenella/go-ansible/pkg/adhoc"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
	"github.com/attapon-th/go-valid-struct/govalidator"
)

type Inventory struct {
	Name string `validate:"required"`
	Host string `validate:"required"`
	Vars map[string]interface{}
}

func main() {
	mac()
	//	makeInventory(inventory())
	//	ansiblePlay()
}

func mac() {
	mac := "AA:BB:CC:DD:EE:FF"
	fmt.Println(regexp.MustCompile(`:`).ReplaceAllString(mac, ""))

	ports := func(n string) string {
		if strings.EqualFold(n, "0") {
			return ""
		}
		return n
	}
	fmt.Println(ports("1-65535"))
	fmt.Println(ports("0"))

}

func ansiblePlay() {
	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		//Connection: "175.45.195.112",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "gotest.ini",
		//ExtraVars: inventory().Vars,
		//	ExtraVars: map[string]interface{}{
		//		"INFRAWARE": inventory().Vars,
		//	},
	}

	playbook := &playbook.AnsiblePlaybookCmd{
		//Playbooks:         []string{"site.yml", "site2.yml"},
		Playbooks:         []string{"test123.yml"},
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		//StdoutCallback:    "json",
		//StdoutCallback: "oneline",
	}

	//log.Println("Command: ", playbook)
	err := playbook.Run(context.TODO())
	if err != nil {
		log.Println("----------------------- ERROR ---------------------------")
		log.Println(err.Error())
		//panic(err)
	} else {
		log.Println("----------------------- SUCCESS ---------------------------")
	}
}

func ansibleAdhoc() {
	ansibleConnectionOptions := &options.AnsibleConnectionOptions{
		//	Connection: "local",
	}

	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		//Inventory:  "127.0.0.1,",
		PlaybookDir: "./",
		ModuleName:  "INFRAWARE",
		ExtraVars:   inventory().Vars,
	}

	adhoc := &adhoc.AnsibleAdhocCmd{
		Pattern:           "all",
		Options:           ansibleAdhocOptions,
		ConnectionOptions: ansibleConnectionOptions,
		//StdoutCallback:    "oneline",
	}

	log.Println("Command: ", adhoc)

	err := adhoc.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}

func inventoryCommonRule() map[string]interface{} {
	return map[string]interface{}{
		"ansible_ssh_user":           "required",
		"ansible_ssh_pass":           "required",
		"ansible_ssh_common_args":    "required",
		"ansible_python_interpreter": "required",
		"t_product":                  "required",
		"t_group":                    "required",
		"t_system":                   "required",
		"t_service":                  "required",
		"t_system_group":             "required",
		"t_system_type":              "required",
		"t_system_id":                "gte=0",
		//"t_flannel_subnet_range":     "",
		//"t_pv_storage_volume":        "",
		//"t_pvc_storage_volume":       "",
	}
}

func inventorySecuregateRule() map[string]interface{} {
	return map[string]interface{}{
		"t_public_ip": "ip",
		//"t_public_vip_ip": "ip",
		//		"t_link_media":    "required",
		//		"t_default_dev":   "required",
	}
}

func inventorySecuregateNicRule() map[string]interface{} {
	return map[string]interface{}{
		"t_in_rx_dev":  "required",
		"t_in_tx_dev":  "required",
		"t_ex_rx_dev":  "required",
		"t_ex_tx_dev":  "required",
		"t_in_rx_mac":  "required",
		"t_ex_rx_mac":  "required",
		"t_in_dest_ip": "required",
		"t_ex_dest_ip": "required",
	}
}

func inventory() *Inventory {
	ini := Inventory{
		Name: "INFRAWARE",
		//Host: "1.1.1.1", // 실패 케이스
		Host: "175.45.195.112", // 성공 케이스
		Vars: map[string]interface{}{
			"ansible_ssh_user":           "root",
			"ansible_ssh_pass":           "hsck@2301",
			"ansible_ssh_common_args":    "\"-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null\"",
			"ansible_python_interpreter": "/usr/bin/python3",
			// Product
			"t_product":      "SG",
			"t_service":      "FILE",
			"t_system_group": "SG0",
			"t_system_type":  "SGIN",
			"t_system_id":    0,
			"t_public_ip":    "1.1.1.1",
		},
	}

	fmt.Println(govalidator.GetValidate().ValidateMap(ini.Vars, inventorySecuregateRule()))
	return &ini
}

func makeInventory(ini *Inventory) {

	iniSlice := []string{}
	iniSlice = append(iniSlice, "["+ini.Name+"]")
	iniSlice = append(iniSlice, ini.Host)
	iniSlice = append(iniSlice, "")

	iniSlice = append(iniSlice, "["+ini.Name+":vars]")
	for key, val := range ini.Vars {
		//iniSlice = append(iniSlice, key+"="+val.(string))
		iniSlice = append(iniSlice, fmt.Sprintf("%v=%v", key, val))
	}

	iniStr := strings.Join(iniSlice, "\n")
	//fmt.Printf("%s", iniStr)

	iniFile, _ := os.Create("gotest.ini")
	defer iniFile.Close()

	iniFile.WriteString(iniStr)
}
