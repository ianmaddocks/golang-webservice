PROJECT?=microservice2
PORT?=8000

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
GOOS?=linux
GOARCH?=amd64
APP?=${PROJECT}
CONTAINER_IMAGE?=docker.io/ianmaddocks/${APP}


clean:
	@echo "clean..."
	rm -f ${APP}

build: clean
	@echo "build..."
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ${APP}

container: build
	@echo "container.."
	docker image rm -f $(CONTAINER_IMAGE):$(RELEASE) || true
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

run: container
	@echo "run..."
	docker stop $(APP):$(RELEASE) || true && docker rm $(APP):$(RELEASE) || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		${CONTAINER_IMAGE}:$(RELEASE)

test:
	@echo "test..."
	go test -v -race ./...

push: container
	@echo "push..."
	docker push $(CONTAINER_IMAGE):$(RELEASE)

minikube: push
	@echo "minikube..."
	for t in $(shell find ./Projects/ci-cd_project/microservice2 -type f -name "*.yaml"); do \
        cat $$t | \
        	gsed -E "s/\{\{(\s*)\.Release(\s*)\}\}/$(RELEASE)/g" | \
        	gsed -E "s/\{\{(\s*)\.ServiceName(\s*)\}\}/$(APP)/g"; \
        echo ---; \
    done > tmp.yaml
	kubectl apply -f tmp.yaml
