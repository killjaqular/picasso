SOLUTION_NAME = "picasso"
REGISTRY = $(SOLUTION_NAME)

# Building recipes
.PHONY: all
all: database server

.PHONY: database
database:
	docker rmi -f $(REGISTRY)/database:latest
	docker build -t $(REGISTRY)/database:latest\
		--no-cache\
		database/

.PHONY: server
server:
	docker rmi -f $(REGISTRY)/server:latest
	docker build -t $(REGISTRY)/server:latest\
		--no-cache\
		server/

# Running recipes
.PHONY: run
run:
	docker compose up

.PHONY: down
stop:
	docker compose down

.PHONY: stop
stop:
	docker compose stop

# Cleaning recipes
.PHONY: clean
clean: stop down
	docker rmi -f\
		$(REGISTRY)/database:latest\
		$(REGISTRY)/server:latest
