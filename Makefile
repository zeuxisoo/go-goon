PACK_ROOT = "release"
PACK_GOON = "release/goon"

NOW = $(shell date -u '+%Y%m%d%I%M%S')

usage:
	@echo "make env"
	@echo "make build"
	@echo "make clean"
	@echo "make pack"

env:
	@glide install

build:
	@go build -o go-goon

build-windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o go-goon.exe

build-freebsd:
	@CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o go-goon.go

build-linix:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-goon.go

clean:
	@rm -rf go-goon
	@rm -rf go-goon.exe
	@rm -rf $(PACK_ROOT)

pack: build
	@rm -rf $(PACK_GOON)
	@mkdir -p $(PACK_GOON)
	@cp -r go-goon README.md conf $(PACK_GOON)
	@rm -rf $(PACK_GOON)/conf/test-*.ini
	@mv $(PACK_GOON)/conf/app.ini.example $(PACK_GOON)/conf/app.ini
	@cd $(PACK_ROOT) && zip -r go-goon.$(NOW).zip "goon"
