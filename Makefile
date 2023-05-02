PACKAGE_LIST := $(shell go list ./...)
VERSION := 0.1.2
NAME := urleap
DIST := $(NAME)-$(VERSION)


urakil: coverage.out
go build -o urleap $(PACKAGE_LIST)

coverage.out:
go test -covermode=count -coverprofile=coverage.out $(PACKAGE_LIST)

test:
        go test -covermode=count -coverprofile=coverage.out $(PACKAGE_LIST)
clean:
        rm -f urakil