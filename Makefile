API=api/raven
RAVEN=api/raven/raven
LOCAL_CONFIG=local/local.yml

all: build

build: .build-api

install: .install-api

dep: .dep-api

run-api: build
	$(RAVEN) run $(LOCAL_CONFIG)

docker-build: build
	doriath -x ravenTag=`$(RAVEN) version` build

clean: .clean-api

.build-api:
	cd $(API) && go build

.install-api:
	cd $(API) && go install

.dep-api:
	cd $(API) && dep ensure

.clean-api:
	cd $(API) && go clean
