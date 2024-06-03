makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list bootstrap init build

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

generate: purge
	go generate ./...

purge:
	rm -rf file_extention/*.gen.go
	rm -rf file_extention/*.enum.go
	rm -rf mime_type/*.gen.go	
	rm -rf mime_type/*.enum.go

test: 
	go test ./...

commit:
	git add . || true
	git commit -m "$(m)" || true
	git push origin master

tag:
	git tag -fa $(tag) -m "$(tag)"
	git push -f origin $(tag)

tag-list:
	git tag -l | sort -V

publish: test
	@if ack replace go.mod ;then echo 'Remove the "replace" line from the go.mod file'; exit 1; fi
	make commit m=$(m)
	make tag tag=$(m)
