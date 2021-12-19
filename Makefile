.PHONY: run
run:
	go1.18beta1 run .

.PHONY: tutorial
tutorial:
	go1.18beta1 run ./tutorial/main.go

.PHONY: yaml
yaml:
	go1.18beta1 run ./yaml/main.go
