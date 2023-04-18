PACKAGE_LIST := $(shell go list ./...)
urakil:
        go build -o urakil $(PACKAGE_LIST)
test:
        go test $(PACKAGE_LIST)
clean:
        rm -f urakil
