.PHONY: help test exercise

# Variables
NAME ?=

# ANSI color codes
BOLD := \033[1m
RESET := \033[0m
GREEN := \033[32m
RED := \033[31m
BLUE := \033[34m

help: ## Show this help message	
	@awk 'BEGIN {FS = ":.*##"; } /^[a-zA-Z_-]+:.*?##/ { printf "$(BOLD)$(BLUE)%-10s$(RESET) %s\n", $$1, $$2 }' $(MAKEFILE_LIST) | sort

test: ## Run unit tests (use NAME=exercisename to target a specific exercise)
	@echo "$(BLUE)> Running tests...$(RESET)"
	@if [ -z "$(NAME)" ]; then \
		go test ./exercises/...; \
	else \
		go test ./exercises/$(NAME); \
	fi

exercise: ## Create a new exercise (use NAME=exercisename)
	@if [ -z "$(NAME)" ]; then \
		echo "$(BOLD)$(RED)Error:$(RESET) NAME is required. Usage: make exercise NAME=exercisename"; \
		exit 1; \
	fi
	@mkdir -p exercises/$(NAME)
	@sed "s/{{NAME}}/$(NAME)/g" stubs/exercise/main.go.stub > exercises/$(NAME)/main.go
	@sed "s/{{NAME}}/$(NAME)/g" stubs/exercise/main_test.go.stub > exercises/$(NAME)/main_test.go
	@echo "$(GREEN)✔ Exercise '$(NAME)' created successfully in exercises/$(NAME)/$(RESET)"

list: ## List all exercises
	@echo "$(BLUE)> Exercises:$(RESET)"
	@ls exercises | sort