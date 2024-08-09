proto:
	protoc --go_opt=module=protocol-buffer-go  --go_out=. ./proto/*.proto


.PHONY: proto