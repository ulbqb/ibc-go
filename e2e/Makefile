RLY_VERSION := 2.3.1
ARCH := $(shell dpkg --print-architecture)

install-relayer:
	@curl -L https://github.com/cosmos/relayer/releases/download/v$(RLY_VERSION)/Cosmos.Relayer_$(RLY_VERSION)_linux_$(ARCH).tar.gz --output relayer.tar.gz
	@tar -xf relayer.tar.gz  --transform 's!^[^/]*!relayer_archive!'
	@mv relayer_archive/rly /usr/bin
	@rly version
	@rm -rf relayer.tar.gz relayer_archive

install-simapp:
	@cd ../ && \
	make install
	@simd version --long

test-e2e:
	@./e2e

.PHONY: install-relayer install-simapp test-e2e
