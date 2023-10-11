package main

// JSON 데이터를 구조체로 매핑하는 구조체 정의
type JSONData struct {
	Label   string    `json:"label"`
	Host    string    `json:"host"`
	Time    string    `json:"time"`
	CPU     []CPU     `json:"cpu"`
	Memory  []Memory  `json:"memory"`
	Disk    []Disk    `json:"disk"`
	Traffic []Traffic `json:"traffic"`
	Process []Process `json:"process"`
}

type CPU struct {
	Time   string `json:"time"`
	Total  int    `json:"total"`
	Idle   int    `json:"idle"`
	System int    `json:"system"`
	Wait   int    `json:"wait"`
	User   int    `json:"user"`
}

type Memory struct {
	Time      string `json:"time"`
	Total     string `json:"total"`
	Available string `json:"available"`
	Use       int    `json:"use"`
	Buffer    int    `json:"buffer"`
	Cache     int    `json:"cache"`
	Free      int    `json:"free"`
}

type Disk struct {
	Mount     string `json:"mount"`
	Available int    `json:"available"`
	Use       int    `json:"use"`
	Total     string `json:"total"`
}

type Traffic struct {
	Name string `json:"name"`
	Rx   int    `json:"rx"`
	Tx   int    `json:"tx"`
}

type Process struct {
	Name   string `json:"name"`
	Count  int    `json:"count"`
	Status string `json:"status"`
}
