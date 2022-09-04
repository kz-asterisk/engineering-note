.DEFAULT_GOAL := help
POETRY := poetry
.PHYONY: setup lint format test deploy help

setup:	## Setup dependencies
				$(POETRY) install

lint: 	## Check if Python files are proper
				$(POETRY) run autoflake --check -ir --remove-all-unused-imports --remove-unused-variables --remove-duplicate-keys .
				$(POETRY) run isort --check --diff .
				$(POETRY) run black --check --diff .

format:	## Format all Python files
				$(POETRY) run autoflake -ir --remove-all-unused-imports --remove-unused-variables --remove-duplicate-keys .
				$(POETRY) run isort .
				$(POETRY) run black .

run:		## Start API Server on LocalEnvironmet
				$(POETRY) run uvicorn main:app --reload

deploy:	## Deploy Python code as XXXX
				@echo "Please Fix it🐠"

help:		## Show this help
				@echo 'Available targets'
		    @grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[35m%-30s\033[0m %s\n", $$1, $$2}'
