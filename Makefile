.PHONY: default fmt create_server linter build

GIT_COMMIT=`git rev-parse HEAD`
GIT_BRANCH=`git rev-parse --abbrev-ref HEAD`
GIT_TAG=`git describe --tags $(git rev-list --tags --max-count=1)`
BUILD_DATE=`date +%FT%T%z`
LDFLAGS=-X main.gitCommit=${GIT_COMMIT} -X main.gitBranch=${GIT_BRANCH} -X main.gitTag=${GIT_TAG} -X main.gitDate=${BUILD_DATE}

# SWAGGER_ARG=-g server.go -d ./internal/server
LINT_ARG=-c .golangci.yaml --timeout 4m

default:
	$(info ************ COMMAND NOT SELECT ************)

lint:
	$(info ************ RUN LINTER ************)
	golangci-lint run ${LINT_ARG}

lint-fix:
	$(info ************ RUN LINTER WITH FIX ************)
	golangci-lint run ${LINT_ARG} --fix

fmt:
	$(info ************ RUN FROMATING ************)
	go fmt ./...
	gofumpt -l -w .

build:
	$(info ************ BUILD TO ./build/app ************)
	CGO_ENABLED=0 GOOS=linux go build  -ldflags "${LDFLAGS}" -a -installsuffix cgo -o=micro ./main.go

mod/download:
	$(info ************ MOD DOWNLOAD ************)
	go mod download

mod/vendor:
	$(info ************ MOD VENDOR ************)
	go mod vendor
	
mod/tidy:
	$(info ************ MOD TIDY ************)
	go mod tidy

test:
	$(info ************ RUN TESTS ************)
	go test -coverpkg=./...  -coverprofile=profile.cov  ./... ;\
	go tool cover -func profile.cov ;\
	rm profile.cov

generate:
	$(info ************ GENERATE MOCKS ************)
	go generate -v ./...

run:
	$(info ************ RUN ************)
	./micro

version:
	$(info ************ VERSION ************)
	./micro --version