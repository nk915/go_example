package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configs struct {
	Manager ManagerConfig `yaml:"Manager"`
}

type ManagerConfig struct {
	Role              string   `yaml:"role"`
	ServerName        string   `mapstructure:"server-name"`
	EthernetInterface string   `mapstructure:"ethernet-interface"`
	RecoverMode       string   `mapstructure:"recover-mode"`
	ProcessScanTime   int      `mapstructure:"process-scan-time"`
	ProcessList       []string `mapstructure:"process-list"`
}

type ProcessConfig struct {
	Running   string        `yaml:"running"`
	Mode      string        `yaml:"mode"`
	Path      string        `yaml:"path"`
	Heartbeat HeartbeatInfo `yaml:"heartbeat"`
}

type HeartbeatInfo struct {
	Use string `mapstructure:"use"`
	Url string `mapstructure:"url"`
}

func main() {
	viper.SetConfigName("service")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println(viper.Get("Manager"))

	/* 전체 파일 내용을 찾으며 언마샬 */
	cfg := Configs{}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	fmt.Println("___ MANAGER (ALL) ___")
	fmt.Println(cfg)

	/* 특정 필드를 찾아 언마샬 */
	var managerConfig = ManagerConfig{}
	if err := viper.UnmarshalKey("Manager", &managerConfig); err != nil {
		panic(err)
	}
	fmt.Println("___ MANAGER (KEY) ___")
	fmt.Println(managerConfig)

	var processConfigList []ProcessConfig
	for _, processName := range managerConfig.ProcessList {
		var processConfig ProcessConfig
		if err := viper.UnmarshalKey(processName, &processConfig); err != nil {
			panic(err)
		}
		// fmt.Println("___ PROCESS ___")
		// fmt.Println(processConfig)
		processConfigList = append(processConfigList, processConfig)
	}
	fmt.Println("___ PROCESS LIST ___")
	fmt.Println(processConfigList)
}
