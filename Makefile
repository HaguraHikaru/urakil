PACKAGE_LIST := $(shell go list ./...)
VERSION := 0.1.5
NAME := urakil
DIST := $(NAME)-$(VERSION)

urakil: coverage.out cmd/urakil/main.go *.go
	go build -o urakil cmd/$(NAME)/main.go cmd/$(NAME)/generate_completion.go

coverage.out: cmd/$(NAME)/main_test.go
	go test -covermode=count \
		-coverprofile=coverage.out $(PACKAGE_LIST)

docker: urakil
#	docker build -t ghcr.io/HaguraHikaru/urakil:$(VERSION) -t ghcr.io/HaguraHikaru/urakil:latest .
	docker buildx build -t ghcr.io/HaguraHikaru/urakil:$(VERSION) \
		-t ghcr.io/HaguraHikaru/urakil:latest --platform=linux/arm64/v8,linux/amd64 --push .

# refer from https://pod.hatenablog.com/entry/2017/06/13/150342
define _createDist
	mkdir -p dist/$(1)_$(2)/$(DIST)
	GOOS=$1 GOARCH=$2 go build -o dist/$(1)_$(2)/$(DIST)/$(NAME)$(3) cmd/$(NAME)/main.go cmd/$(NAME)/generate_completion.go
	cp -r README.md LICENSE dist/$(1)_$(2)/$(DIST)
#	cp -r docs/public dist/$(1)_$(2)/$(DIST)/docs
	tar cfz dist/$(DIST)_$(1)_$(2).tar.gz -C dist/$(1)_$(2) $(DIST)
endef

dist: urakil
	@$(call _createDist,darwin,amd64,)
	@$(call _createDist,darwin,arm64,)
	@$(call _createDist,windows,amd64,.exe)
	@$(call _createDist,windows,arm64,.exe)
	@$(call _createDist,linux,amd64,)
	@$(call _createDist,linux,arm64,)

distclean: clean
	rm -rf dist

clean:
	rm -f urakil coverage.out