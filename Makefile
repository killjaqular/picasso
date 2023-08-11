SOLUTION_NAME = "picasso"
REGISTRY = $(SOLUTION_NAME)

# Building recipes

.PHONY: all
all: database server

.PHONY: database
database:
	docker build -t $(REGISTRY)/database:latest\
		database/

.PHONY: server
server:
	go build server/main.go
	docker build -t $(REGISTRY)/server:latest\
		server/

# Running recipes

.PHONY: run
run:
	docker compose up

.PHONY: stop
stop:
	docker compose stop

# Cleaning recipes

.PHONY: clean
clean: stop
	docker rmi -f\
		$(REGISTRY)/database:latest\
		$(REGISTRY)/server:latest
