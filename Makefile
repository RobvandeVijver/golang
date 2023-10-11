.DEFAULT_GOAL := help

.PHONY: help
help: ## Show the available commands
	@printf "\033[33mUsage:\033[0m\n  make [target] [arg=\"val\"...]\n\n\033[33mTargets:\033[0m\n"
	@grep -E '^[-a-zA-Z0-9_\.\/]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[32m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the docker image
	@printf "\033[32mBuilding docker image...\033[0m\n"
	@docker compose build

.PHONY: run
run: ## Run the docker image
	@printf "\033[32mRunning docker image...\033[0m\n"
	@docker compose up -d

.PHONY: stop
stop: ## Stop the docker image
	@printf "\033[32mStopping docker image...\033[0m\n"
	@docker compose stop

.PHONY: clean
clean: ## Clean the docker image
	@printf "\033[32mCleaning docker image...\033[0m\n"
	@docker compose down --rmi all --volumes --remove-orphans

.PHONY: sh
sh: ## Run a shell in the docker image
	@printf "\033[32mRunning shell in docker image...\033[0m\n"
	@docker compose exec app sh

.PHONY: logs
logs: ## Show the logs of the docker image
	@printf "\033[32mShowing logs of docker image...\033[0m\n"
	@docker compose logs -f app --tail 5


.PHONY: usecase 
usecase: ## part 1
	@printf "\033[32mRunning go application in docker image...\033[0m\n"
	@docker compose exec app ./bin/main details -imdbid tt0034583

	@printf "\033[32mRunning list command...\033[0m\n"
	@docker compose exec app ./bin/main list

	@printf "\033[32mRunning details command...\033[0m\n"
	@docker compose exec app ./bin/main details -imdbid tt0034583

	@printf "\033[32mRunning add command...\033[0m\n"
	@docker compose exec app ./bin/main add -imdbid tt10872600 -title "Spider-Man: No Way Home" -rating 8.3 -year 2021

	@printf "\033[32mRunning delete command...\033[0m\n"
	@docker compose exec app ./bin/main delete -imdbid tt0058150

.PHONY: usecase2
usecase2: ## part 2
	 @printf "\033[32mRunning /movies commands...\033[0m\n"
	 @docker compose exec app curl -S -s localhost:8090/movies

	 @printf "\033[32mRunning /movies/tt0034583 commands...\033[0m\n"
	 @docker compose exec app curl -S -s localhost:8090/movies/tt0034583

	 @printf "\033[32mRunning /movies/tt0034583 commands...\033[0m\n"
	 @docker compose exec app curl -S -s -H "Content-Type: application/json" -X POST -d '{"imdb_id": "tt0368226", "title": "The Room", "rating": 3.7, "year": 2003}' localhost:8090/movies

	@printf "\033[32mRunning localhost:8090/movies/tt0058150 | head -n 1 commands...\033[0m\n"
	@docker compose exec app curl -i -s -X DELETE localhost:8090/movies/tt0058150 | head -n 1

.PHONY: usecase3
usecase3: ## part 3
	 @printf "\033[32mRunning *text* commands...\033[0m\n"
	 @docker compose exec app curl -S -s localhost:8090/movies

