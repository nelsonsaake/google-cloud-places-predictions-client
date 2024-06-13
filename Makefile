.PHONY: run get ncommit

NELSONSAAKE_GO_VERSION = v0.0.22

run:
	cls
	go run .

get:
	cls
	go get github.com/nelsonsaake/go@$(NELSONSAAKE_GO_VERSION)

ncommit:
	cls
	git add .
	git commit -m "ncommit"
	git push origin main