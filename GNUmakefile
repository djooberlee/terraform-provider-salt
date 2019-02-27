LDFLAGS += -X main.version=$$(git describe --always --abbrev=40 --dirty)

# default  args for tests
TEST_ARGS_DEF := -covermode=count -coverprofile=profile.cov

default: build

terraform-provider-salt:
	go build -ldflags "${LDFLAGS}"

build: fmt-check lint-check vet-check terraform-provider-salt

install:
	go install -ldflags "${LDFLAGS}"

# unit tests
# usage:
# - run all the unit tests: make test
# - run some particular test: make test TEST_ARGS="-run TestAccSaltDomain_Cpu"
test:
	go test -v $(TEST_ARGS_DEF) $(TEST_ARGS) ./salt

# acceptance tests
# usage:
#
# - run all the acceptance tests:
#   make testacc
#
# - run some particular test:
#   make testacc TEST_ARGS="-run TestAccSaltDomain_Cpu"
#
# - run all the network test with a verbose loglevel:
#   TF_LOG=DEBUG make testacc TEST_ARGS="-run TestAccSaltNet*"
#
testacc:
	./travis/run-tests-acceptance $(TEST_ARGS)

vet-check:
	go vet ./salt

lint-check:
	go run golang.org/x/lint/golint -set_exit_status ./salt .

fmt-check:
	go fmt ./salt .

tf-check:
	terraform fmt -write=false -check=true -diff=true examples/

clean:
	rm -f terraform-provider-salt

cleanup:
	./travis/cleanup.sh

.PHONY: build install test testacc vet-check fmt-check lint-check
