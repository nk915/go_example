package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

// https://github.com/opencopilot/consulkvjson

func main() {

	//config.Address = "192.168.1.45:8501" // Consul 서버 주소

	config := config_120()
	fileName := "backup_120.json"

	if data, err := ConsulToMap(config, "/"); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println(data)
		save(fileName, data)
	}

	//	fileName := "write.json"
	//	rdata, err := load2(fileName)
	//	if err != nil {
	//		fmt.Println("err:", err)
	//		return
	//	} else {
	//		fmt.Printf("%+v", rdata)
	//	}
	//
	//	config := config_localhost()
	//	if _, err := MapToConsul(config, rdata); err != nil {
	//		fmt.Println("err:", err)
	//		return
	//	}
	fmt.Println("success")

}

func config_localhost() *consulapi.Config {
	config := consulapi.DefaultConfig()
	config.Address = "localhost:8501" // Consul 서버 주소
	//	config.Scheme = "https"
	//	config.TLSConfig = consulapi.TLSConfig{
	//		Address:            config.Address,
	//		CAFile:             "C:/workspace/__GitHub/__Nk915/go_example/externel-library/consul/watching/consul-agent-ca.pem",
	//		CertFile:           "C:/workspace/__GitHub/__Nk915/go_example/externel-library/consul/watching/ex-server-consul-0.pem",
	//		KeyFile:            "C:/workspace/__GitHub/__Nk915/go_example/externel-library/consul/watching/ex-server-consul-0-key.pem",
	//		InsecureSkipVerify: false,
	//	}

	return config
}

func config_120() *consulapi.Config {
	config := consulapi.DefaultConfig()
	config.Address = "192.168.20.120:8501" // Consul 서버 주소
	config.Scheme = "https"
	config.TLSConfig = consulapi.TLSConfig{
		Address:            config.Address,
		CAFile:             "C:/workspace/__GitHub/__Nk915/go_example/externel-library/consul/watching/consul-agent-ca.pem",
		CertFile:           "C:/workspace/__GitHub/__Nk915/go_example/externel-library/consul/watching/ex-server-consul-0.pem",
		KeyFile:            "C:/workspace/__GitHub/__Nk915/go_example/externel-library/consul/watching/ex-server-consul-0-key.pem",
		InsecureSkipVerify: false,
	}

	return config
}

func config_110() *consulapi.Config {
	config := consulapi.DefaultConfig()
	config.Address = "192.168.20.110:8501" // Consul 서버 주소
	config.Scheme = "https"
	config.TLSConfig = consulapi.TLSConfig{
		Address:            config.Address,
		CAFile:             "C:/workspace/__GitHub/__Nk915/go_example/externel-library/consul/watching/in/consul-agent-ca.pem",
		CertFile:           "C:/workspace/__GitHub/__Nk915/go_example/externel-library/consul/watching/in/in-server-consul-0.pem",
		KeyFile:            "C:/workspace/__GitHub/__Nk915/go_example/externel-library/consul/watching/in/in-server-consul-0-key.pem",
		InsecureSkipVerify: false,
	}

	return config
}

func load(fromFilePath string) (map[string]string, error) {
	rFile, err := ioutil.ReadFile(fromFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// JSON 데이터를 map[string]interface{}로 언마샬링
	var jsonData map[string]interface{}
	err = json.Unmarshal(rFile, &jsonData)
	if err != nil {
		log.Fatal(err)
	}

	// 결과를 저장할 map 생성
	result := make(map[string]string)

	// 최상위 키를 유지하면서, 나머지 하위 구조를 JSON 문자열로 변환
	for key, value := range jsonData {
		// 하위 구조를 들여쓰기가 적용된 JSON 문자열로 변환
		bytes, err := json.MarshalIndent(value, "", "  ") // 여기가 변경된 부분입니다.
		if err != nil {
			log.Fatal(err)
		}
		result[key] = string(bytes)
		//	switch valueTyped := value.(type) {
		//	case map[string]interface{}:
		//		// 하위 구조를 JSON 문자열로 변환
		//		bytes, err := json.Marshal(valueTyped)
		//		if err != nil {
		//			log.Fatal(err)
		//		}
		//		result[key] = string(bytes)
		//	default:
		//		// 비정상적인 경우(예상 외의 구조), 에러 처리
		//		log.Fatalf("Unexpected structure under key %s", key)
		//	}
	}

	return result, nil
}

// ConsulToMap takes a consul config and a path offset
// Connects to consul "key/value".
// Reads all (i.e. "recurse") {k, v} pairs under the path offset
// into a map[string]string preserving path hierarchy in map keys: i.e. {"universe/answers/main": "42"}
func ConsulToMap(consulSpec *consulapi.Config, offset string, keysWithOffset ...bool) (map[string]string, error) {

	consul, err := consulapi.NewClient(consulSpec)
	if err != nil {
		return nil, err
	}

	kv := consul.KV()

	config := make(map[string]string)

	kvps, _, err := kv.List(offset, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch k/v pairs from consul: %+v, path offset: %s. due to %v", consulSpec, offset, err)
	}

	withOffset := true
	if len(keysWithOffset) > 0 {
		withOffset = keysWithOffset[0]
	}

	for _, kvp := range kvps {
		if val := kvp.Value; val != nil {
			k := kvp.Key
			if !withOffset {
				k = strings.Split(kvp.Key, offset)[1]
			}
			config[k] = string(val[:])
		}
	}

	log.Printf("read consul map at offset: /%s\n", offset)

	//	excludeList := []string{"password", "secret", "auth"}
	//	for k, _ := range config {
	//		_, found := func(slice []string, key string) (int, bool) {
	//			for i, item := range slice {
	//				search := strings.Contains(key, item)
	//				if search != false {
	//					return i, true
	//				}
	//			}
	//			return -1, false
	//		}(excludeList, k)
	//
	//		if !found {
	//			//		fmt.Printf("read consul map entry: {:%s, %s }\n", k, v)
	//		} else {
	//			//		fmt.Printf("read consul map entry: {:%s, %s }\n", k, "*******")
	//		}
	//	}

	return config, nil
}

// MapToConsul takes a consul config and a map[string]string
// Connects to consul "key/value".
// Walks over a given map and "PUT"s its etries to consul
// respecting path hierarchy encoded in keys: i.e. {"universe/answer/main": 42}.
// Returns a total time.Duration of all the "PUT" operations
func MapToConsul(consulSpec *consulapi.Config, config map[string]string) (time.Duration, error) {

	consul, err := consulapi.NewClient(consulSpec)
	if err != nil {
		return 0, err
	}

	kv := consul.KV()

	var duration time.Duration

	for k, v := range config {
		took, err := kv.Put(&consulapi.KVPair{Key: k, Value: []byte(v)}, nil)
		if err != nil {
			return 0, fmt.Errorf("could not put a key, value: {%s, %s} to consul %+v due to %v", k, v, consulSpec, err)
		}
		duration += took.RequestTime
	}

	return duration, nil
}

func save(toFilePath string, from map[string]string) error {
	// map을 JSON으로 변환
	jsonData, err := json.Marshal(from)
	if err != nil {
		log.Fatal(err)
	}

	// JSON 데이터를 파일로 저장
	err = ioutil.WriteFile(toFilePath, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func load2(fromFilePath string) (map[string]string, error) {
	rFile, err := ioutil.ReadFile(fromFilePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println(string(rFile))

	rMap := make(map[string]string)
	err = json.Unmarshal(rFile, &rMap)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return rMap, err
}
