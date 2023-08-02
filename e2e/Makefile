RLY_VERSION := 2.3.1
CSMS_VERSION := 4.3.1
ARCH := $(shell dpkg --print-architecture)
PATH := $(shell pwd)/bin:$(PATH)

bin/simd: 
	@mkdir -p bin && \
	cd ../ && \
	make build && \
	mv build/simd e2e/bin
	@simd version --long

bin/csmssimd:
	@mkdir -p bin/cosmos && \
	cd bin/cosmos && \
	curl -L https://github.com/cosmos/ibc-go/releases/download/v$(CSMS_VERSION)/ibc-go_simd_v$(CSMS_VERSION)_linux_$(ARCH).tar.gz --output cosmos_simd.tar.gz && \
	tar -xf cosmos_simd.tar.gz && \
	mv simd ../csmssimd
	@csmssimd version --long --log_level=info

bin/rly:
	@mkdir -p bin/relayer && \
	cd bin/relayer && \
	curl -L https://github.com/cosmos/relayer/releases/download/v$(RLY_VERSION)/Cosmos.Relayer_$(RLY_VERSION)_linux_$(ARCH).tar.gz --output relayer.tar.gz && \
	tar -xf relayer.tar.gz  --transform 's!^[^/]*!relayer_archive!' && \
	mv relayer_archive/rly ../
	@rly version

test-fnsa-fnsa: bin/rly bin/simd
	@./e2e fnsa-fnsa fnsa-0 simd 438 fnsa-1 simd 438 "skip" 

test-fnsa-csms: bin/rly bin/simd bin/csmssimd
	@./e2e fnsa-csms fnsa-0 simd 438 csms-1 csmssimd 118 "skip" 

clean-all:
	@rm -rf bin data

clean-data:
	@rm -rf data

.PHONY: test-fnsa-fnsa test-fnsa-csms clean-all clean-data
