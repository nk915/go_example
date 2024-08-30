package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func printBody(body io.Reader) string {
	read, err := io.ReadAll(body)
	if err != nil {
		return fmt.Sprintf("Failed to read response body: %s", err.Error())
	}
	return string(read)
}

func authenticate(harborUrl, username, password string) bool {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	req, _ := http.NewRequest("GET", harborUrl+"/users/current", nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Authentication failed:", err)
		return false
	}
	defer resp.Body.Close()

	fmt.Println("response: ", printBody(resp.Body))

	return resp.StatusCode == 200
}

func getArtifacts(harborUrl, projectName, repositoryName, username, password string) []map[string]interface{} {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/projects/%s/repositories/%s/artifacts", harborUrl, projectName, repositoryName), nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to retrieve artifacts:", err)
		return nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response: ", string(body))

	var artifacts []map[string]interface{}
	json.Unmarshal(body, &artifacts)
	return artifacts
}

func checkNotarySignatures(notaryUrl, repositoryName string, artifacts []map[string]interface{}, username, password string) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	for _, artifact := range artifacts {
		digest := artifact["digest"].(string)
		req, _ := http.NewRequest("GET", fmt.Sprintf("%s/v2/%s/_trust/tuf/%s/metadata", notaryUrl, repositoryName, digest), nil)
		req.SetBasicAuth(username, password)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Failed to check notary signature:", err)
			continue
		}
		if resp.StatusCode == 200 {
			fmt.Printf("Digest %s is signed\n", digest)
		} else {
			fmt.Printf("Digest %s is not signed\n", digest)
		}
	}
}

func main() {
	harbor := "https://<harbor>:3443"
	username := "<user-name>"
	password := "<password>"
	harborUrl := harbor + "/api/v2.0"

	if authenticate(harborUrl, username, password) {
		fmt.Println("Authentication successful")
	} else {
		fmt.Println("Authentication failed")
		return
	}

	projectName := "web"
	repositoryName := "securegate%252Fweblink"
	artifacts := getArtifacts(harborUrl, projectName, repositoryName, username, password)
	if artifacts == nil {
		return
	}

	notaryUrl := harbor + "/notary"
	checkNotarySignatures(notaryUrl, repositoryName, artifacts, username, password)
}
