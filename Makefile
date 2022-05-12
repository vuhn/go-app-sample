GO_BUILD_IMAGE_VER=v1.0.0

start-api:
	GO111MODULE=on; cd src/cmd/api; go run .

mocks:
	cd src/pkg; mockery --all
	cd src/infrastructure/repository; mockery --all
	cd src/service; mockery --all

unittest:
	GO111MODULE=on; cd src; go test --cover ./...

run:
	docker-compose up -d

docker-build-ci:
	docker build -f ./.docker/DockerfileBuild -t vuhn07/gobuild:${GO_BUILD_IMAGE_VER} .

docker-push-ci:
	docker push vuhn07/gobuild:${GO_BUILD_IMAGE_VER}
