GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
DOCKER_IMAGE_NAME=liu578101804/up_to_qiniu:1.0

.PHONY: dev
dev:
	$(GOBUILD) -o ./build/app -v && ./build/app

.PHONY: build_linux
build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/app_linux -v

.PHONY: build_docker
build_docker: build_linux
	docker build -t $(DOCKER_IMAGE_NAME) .

.PHONY: up_docker
up_docker: build_docker
	docker push $(DOCKER_IMAGE_NAME)
