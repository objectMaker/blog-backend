SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

run:
	go run main.go

build:
	docker build \
	-t blog-backend:1.0 \
	.
	
up:
	docker-compose up