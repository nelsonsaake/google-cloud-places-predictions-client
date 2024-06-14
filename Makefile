.PHONY: run get ncommit

NELSONSAAKE_GO_VERSION = v0.0.38
ARGS = $(filter-out $@,$(MAKECMDGOALS))

run:
	cls
	go run . $(ARGS)

get:
	cls
	go get github.com/nelsonsaake/go@$(NELSONSAAKE_GO_VERSION)

ncommit:
	cls
	git add .
	git commit -m "ncommit"
	git push origin main