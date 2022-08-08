SERVICE=auth_service

build:
	@echo "building ${SERVICE}..."
	@go build -o dist/${SERVICE} main.go

run:
	@echo "running ${SERVICE}..."
	./dist/${SERVICE}

clean:
	@echo "cleaning dist/"
	@rm -r dist/${SERVICE}
