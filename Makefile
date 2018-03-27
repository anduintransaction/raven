BIN=raven
API=api/raven
FRONTEND=frontend
FRONTEND_BUILD=$(FRONTEND)/build
RAVEN=$(API)/$(BIN)
LOCAL_CONFIG=local/local.yml
DIST=dist

OS_MAC=darwin
ARCH_MAC=amd64
OS_LINUX=linux
ARCH_LINUX=amd64

current_dir=$(shell pwd)

all: build

build: .build-api .build-frontend

install: .install-api

dep: .dep-api

run: build
	$(RAVEN) run $(LOCAL_CONFIG) --ui-data=$(FRONTEND)/build

run-api: .build-api
	$(RAVEN) run $(LOCAL_CONFIG)

run-frontend:
	cd frontend && yarn start

dist: .build-api .build-frontend .dist-mac .dist-linux .dist-post

clean: .clean-api .clean-frontend .clean-dist

.build-api:
	cd $(API) && go build

.build-frontend:
	cd $(FRONTEND) && yarn build

.build-dist-api-mac:
	cd $(API) && GOOS=$(OS_MAC) GOARCH=$(ARCH_MAC) go build -o $(current_dir)/$(DIST)/raven/$(BIN)

.build-dist-api-linux:
	cd $(API) && GOOS=$(OS_LINUX) GOARCH=$(ARCH_LINUX) go build -o $(current_dir)/$(DIST)/raven/$(BIN)

.install-api:
	cd $(API) && go install

.dep-api:
	cd $(API) && dep ensure

.dist-prepare:
	rm -rf $(DIST)/raven

.dist-copy-frontend:
	cp -rf $(FRONTEND)/build $(current_dir)/$(DIST)/raven/frontend

.dist-copy-resources:
	cp -rf README.md LICENSE config.sample.yml $(current_dir)/$(DIST)/raven

.dist-gzip-mac:
	cd $(current_dir)/$(DIST) && tar czf raven-$(OS_MAC)-$(ARCH_MAC)-v`../$(RAVEN) version`.tar.gz raven

.dist-gzip-linux:
	cd $(current_dir)/$(DIST) && tar czf raven-$(OS_LINUX)-$(ARCH_LINUX)-v`../$(RAVEN) version`.tar.gz raven

.dist-post:
	rm -rf $(DIST)/raven

.dist-mac: .build-dist-api-mac .dist-copy-frontend .dist-copy-resources .dist-gzip-mac

.dist-linux: .build-dist-api-linux .dist-copy-frontend .dist-copy-resources .dist-gzip-linux

.clean-api:
	cd $(API) && go clean

.clean-frontend:
	rm -rf $(FRONTEND_BUILD)

.clean-dist:
	rm -rf $(DIST)
