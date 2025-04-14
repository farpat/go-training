.PHONY: test help exercise

# PATH can be specified to test a specific exercise
# Example: make test PATH=./exercises/countbits
PATH ?= ./...

# NAME is required for creating a new exercise
# Example: make exercise NAME=example
NAME ?=

## help: Display this help message
help:
	@echo "Available commands:"
	@echo
	@grep -E '^##' $(MAKEFILE_LIST) | sed -e 's/## //' | sort

## test: Run unit tests (use PATH=./exercises/name for a specific exercise)
test:
	@echo "Running tests..."
	@if [ "$(PATH)" = "./..." ]; then \
		for dir in exercises/*/ ; do \
			if [ -d "$$dir" ]; then \
				echo "Testing $${dir}..."; \
				cd "$$dir" && go test && cd ../..; \
			fi \
		done \
	else \
		go test $(PATH); \
	fi

## exercise: Create a new exercise (use NAME=exercisename)
exercise:
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME parameter is required. Usage: make exercise NAME=exercisename"; \
		exit 1; \
	fi
	@echo "Creating new exercise: $(NAME)"
	@mkdir -p exercises/$(NAME)
	@echo "// Package $(NAME) implements the $(NAME) exercise" > exercises/$(NAME)/main.go
	@echo "package $(NAME)" >> exercises/$(NAME)/main.go
	@echo "" >> exercises/$(NAME)/main.go
	@echo "// MainFunction is the main function of the exercise" >> exercises/$(NAME)/main.go
	@echo "// TODO: rename this function and update its signature according to the exercise" >> exercises/$(NAME)/main.go
	@echo "func MainFunction(input int) int {" >> exercises/$(NAME)/main.go
	@echo "	// TODO: implement your solution" >> exercises/$(NAME)/main.go
	@echo "	return input" >> exercises/$(NAME)/main.go
	@echo "}" >> exercises/$(NAME)/main.go
	@echo "" >> exercises/$(NAME)/main.go
	@echo "package $(NAME)" > exercises/$(NAME)/main_test.go
	@echo "" >> exercises/$(NAME)/main_test.go
	@echo "import \"testing\"" >> exercises/$(NAME)/main_test.go
	@echo "" >> exercises/$(NAME)/main_test.go
	@echo "func TestMainFunction(t *testing.T) {" >> exercises/$(NAME)/main_test.go
	@echo "	// TODO: implement your test" >> exercises/$(NAME)/main_test.go
	@echo "	result := MainFunction(42)" >> exercises/$(NAME)/main_test.go
	@echo "	if result != 42 {" >> exercises/$(NAME)/main_test.go
	@echo "		t.Errorf(\"MainFunction(42) = %d; expected 42\", result)" >> exercises/$(NAME)/main_test.go
	@echo "	}" >> exercises/$(NAME)/main_test.go
	@echo "}" >> exercises/$(NAME)/main_test.go
	@echo "" >> exercises/$(NAME)/main_test.go
	@echo "Exercise $(NAME) created successfully in exercises/$(NAME)/"
