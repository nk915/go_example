PROJ=test-integrity
DOCKER_REGISTRY_IP=192.168.20.125
DOCKER_REGISTRY=$(DOCKER_REGISTRY_IP):3443
DOCKER_REGISTRY_PROJECT=$(DOCKER_REGISTRY)/securegate
DOCKER_IMAGE=$(PROJ)

all: build 

clean:
	go clean -testcache
#	rm -rf msa
#	rm -rf agent
	docker rmi $(DOCKER_REGISTRY_PROJECT)/$(PROJ)

build:
	go mod tidy
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o msa main.go 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o agent main.go 


docker:
	docker build --rm -t $(DOCKER_REGISTRY_PROJECT)/$(PROJ) .


docker-push:
	docker push $(DOCKER_REGISTRY_PROJECT)/$(PROJ)