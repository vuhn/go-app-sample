GO_BUILD_IMAGE_VER=v1.0.0
VER=dev
GCP_PROJECT_ID=kubenetes-learning-vuhn

start-api:
	GO111MODULE=on; cd src/cmd/api; go run .

mocks:
	cd src/pkg; mockery --all
	cd src/repository; mockery --all
	cd src/service; mockery --all

unittest:
	GO111MODULE=on; cd src; go test --cover ./...

run:
	docker-compose up -d

docker-build-ci:
	docker build -f ./.docker/DockerfileBuild -t vuhn07/gobuild:${GO_BUILD_IMAGE_VER} .

docker-push-ci:
	docker push vuhn07/gobuild:${GO_BUILD_IMAGE_VER}

ci-deploy-gcp:
	gcloud auth activate-service-account --key-file=$(KEYFILE)
	gcloud config set project ${GCP_PROJECT_ID}
	make deploy-gcp

deploy-gcp:
	cp deployment/app.yaml src/app.yaml
	cat deployment/app_evn_dev.yaml > src/app_env.yaml
	echo "  DB_HOST: $(DB_HOST)" >> src/app_env.yaml
	echo "  DB_PASSWORD: $(DB_PASSWORD)" >> src/app_env.yaml
	echo "  JWT_KEY: $(JWT_KEY)" >> src/app_env.yaml
	cd src; gcloud app deploy \
		--project ${GCP_PROJECT_ID} \
		--version ${VER} \
		--no-promote \
		--quiet \
		./app.yaml
	rm src/app_env.yaml
	rm src/app.yaml
