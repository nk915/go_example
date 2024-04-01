package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func main() {
	// Consul 클라이언트 구성
	config := api.DefaultConfig()
	config.Address = "localhost:8501" // Consul 서버 주소

	//	config.Address = "192.168.20.120:8501" // Consul 서버 주소
	//	config.Scheme = "https"
	//	config.TLSConfig = api.TLSConfig{
	//		Address:            config.Address,
	//		CAFile:             "C:/workspace/__GitHub/__Nk915/go_example/externel-library/consul/watching/consul-agent-ca.pem",
	//		CertFile:           "C:/workspace/__GitHub/__Nk915/go_example/externel-library/consul/watching/ex-server-consul-0.pem",
	//		KeyFile:            "C:/workspace/__GitHub/__Nk915/go_example/externel-library/consul/watching/ex-server-consul-0-key.pem",
	//		InsecureSkipVerify: false,
	//	}

	//	config.Scheme = "https"
	//	config.TLSConfig = api.TLSConfig{
	//		Address: "192.168.20.120:8501",
	//		//	CAFile:             "path/to/ca_cert.pem",     // CA 인증서 파일 경로
	//		//	CertFile:           "path/to/client_cert.pem", // 클라이언트 인증서 파일 경로
	//		//	KeyFile:            "path/to/client_key.pem",  // 클라이언트 키 파일 경로
	//		CAFile:             "C:/workspace/__GitHub/__Securegate_v4.0/meta/example/consul-agent-ca.pem",
	//		CertFile:           "C:/workspace/__GitHub/__Securegate_v4.0/meta/example/ex-server-consul-0.pem",
	//		KeyFile:            "C:/workspace/__GitHub/__Securegate_v4.0/meta/example/ex-server-consul-0-key.pem",
	//		InsecureSkipVerify: false,
	//	}

	consul(config)
	//watchViper(config)
}

func watchViper(config *api.Config) {
	_, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Consul 클라이언트 생성 실패: %v", err)
	}

	// 원격 설정 소스로 Consul 사용 설정
	if err := viper.AddRemoteProvider("consul", config.Address, "namespace/service/sgate/E001/task"); err != nil {
		log.Fatalf("AddRemoteProvider %v", err)
	}
	viper.SetConfigType("json") // 설정 파일의 타입 지정 (예: "json")

	// 초기 원격 설정 읽기
	err = viper.ReadRemoteConfig()
	if err != nil {
		log.Fatalf("원격 설정 읽기 실패: %v", err)
	}

	// 설정 변경 감지 및 적용
	go func() {
		for {
			time.Sleep(time.Second * 5) // 5초마다 설정 변경 감지

			err := viper.WatchRemoteConfig()
			if err != nil {
				log.Printf("원격 설정 감시 중 오류 발생: %v", err)
				continue
			}

			// 설정이 변경되었을 때 실행할 로직
			fmt.Println("새로운 설정 적용됨:", viper.GetString("process"))
		}
	}()

	select {} // 메인 고루틴이 종료되지 않도록 대기
}

func consul(config *api.Config) {
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Consul 클라이언트 생성 오류: %v", err)
	}

	if client == nil {
		log.Fatalf("Consul 클라이언트 : %v", client)
		return
	}

	// KV 예시: 특정 키의 변경을 감시
	kv := client.KV()
	options := &api.QueryOptions{}

	go func() {
		for {
			pair, meta, err := kv.Get("namespace/bb", options)
			if err != nil {
				log.Printf("KV 가져오기 오류: %v", err)
				continue
			}
			fmt.Printf("%+v\n", meta)
			options.WaitIndex = meta.LastIndex // 변경 감지를 위해 LastIndex 업데이트
			fmt.Printf("키: %v, 값: %v\n", pair.Key, string(pair.Value))
			fmt.Printf("\n\n")
		}
	}()

	// 프로그램이 바로 종료되지 않도록 대기
	select {}

}
