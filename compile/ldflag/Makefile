PROJ=myapp
NAMESPACE=custom


all: build

clean:
	rm $(PROJ) 


build:
	go build -ldflags "-X 'main.namespace=$(NAMESPACE)/'" -o $(PROJ) main.go

