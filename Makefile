makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list bootstrap init build

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

process:
	cd gen && go run process.go

purge:
	rm -rf mime_type/enumeration.go
	rm -rf mime-types.txt
	rm -rf file_extention/enumeration.go
	rm -rf file-extentions.txt

build: purge process
	go generate ./...
