DOCKER_REGISTRY_IP=192.168.20.126
DOCKER_REGISTRY=$(DOCKER_REGISTRY_IP):3443
DOCKER_REGISTRY_PROJECT=$(DOCKER_REGISTRY)/core-platform
PROJ=sevice-agent
SERVICE=service
AGENT=agent


#all: build-windows
all: build-linux

clean:
	rm *.exe
	rm $(SERVICE)
	rm $(AGENT)

build-windows:
	go mod tidy
	CGO_ENABLED=0 go build -a -ldflags '-s -w' -o $(SERVICE).exe $(SERVICE)_project/main.go 
	CGO_ENABLED=0 go build -a -ldflags '-s -w' -o $(AGENT).exe $(AGENT)_project/main.go 
#	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o $(PROJ).exe cmd/main.go cmd/consul_init.go

build-linux:
	go mod tidy
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o $(SERVICE) $(SERVICE)_project/main.go 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o $(AGENT) $(AGENT)_project/main.go 

docker:
	docker build --rm -t $(DOCKER_REGISTRY_PROJECT)/$(PROJ) .

docker-push:
	docker push $(DOCKER_REGISTRY_PROJECT)/$(PROJ)