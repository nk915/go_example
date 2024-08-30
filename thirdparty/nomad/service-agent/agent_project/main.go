package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	EXEC_PATH = "./service.exe"
)

func main() {
	log := logrus.New()
	log.Infoln("[agent] Starting A process...")

	// Start and monitor the A process
	err := startAndMonitorProcess(log)
	if err != nil {
		log.Fatalf("[agent] Error: %v", err)
	}
}

func startAndMonitorProcess(log *logrus.Logger) error {
	cmd := exec.Command("go", "run", "service_project/main.go")
	//cmd := exec.Command("./service.exe")

	// Create pipes for stdout and stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdout pipe: %v", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to create stderr pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %v", err)
	}

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			text := scanner.Text()
			log.Infof("[A process] %s" + text)
		}
		if err := scanner.Err(); err != nil {
			log.Errorf("[agent] Error reading stdout: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			text := scanner.Text()
			if strings.Contains(text, "Error") {
				log.Errorf("[A process] %s" + text)
			}
		}
		if err := scanner.Err(); err != nil {
			log.Errorf("[agent] Error reading stderr: %v", err)
		}
	}()

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("command finished with error: %v", err)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	log.Infoln("[agent] A process ended.")
	return nil
}
