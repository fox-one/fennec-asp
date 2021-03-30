COMMIT = $(shell git rev-parse --short HEAD)
VERSION = $(shell git describe --abbrev=0) 

REPOSITORY_PATH = $(shell cat .config.ini)
ENV = $*
GO = go

%-local: GO = GO111MODULE=on go
%-test: GO = GOOS=linux GOARCH=amd64 GO111MODULE=on go
%-prod: GO = GOOS=linux GOARCH=amd64 GO111MODULE=on go

.PHONY: clean
clean-%:	
	@echo "cleaning building caches and configs......................."
	${GO} clean
	rm -f ./fennec
	rm -f ./fennec.${ENV}
	rm -f ./config/config.yaml

sync-%: 
	@echo "sync code and config file..........................."
	git pull
	cp -f ./deploy/config.${ENV}.yaml ./config/config.yaml

build-%: clean-% sync-%
	${GO} build --ldflags "-s -w -X main.version=${VERSION} -X main.commit=${COMMIT}"
	cp ./fennec ./fennec.${ENV}

docker-build-%: build-%
	@echo "repository path -> ${REPOSITORY_PATH}"
	docker build -t ${REPOSITORY_PATH}/fennec-${ENV}:${VERSION} . 

.PHONY: aws-login
aws-login:
	$(shell aws ecr get-login --no-include-email --region ap-northeast-1)

deploy-%: docker-build-%
	docker push ${REPOSITORY_PATH}/compound-${ENV}:${VERSION}