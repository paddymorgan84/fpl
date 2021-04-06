.PHONY: explain
explain:
	### Welcome
	#
	# FFFFFFFFFFFFFFFFFFFFFFPPPPPPPPPPPPPPPPP   LLLLLLLLLLL
	# F::::::::::::::::::::FP::::::::::::::::P  L:::::::::L
	# F::::::::::::::::::::FP::::::PPPPPP:::::P L:::::::::L
	# FF::::::FFFFFFFFF::::FPP:::::P     P:::::PLL:::::::LL
	#   F:::::F       FFFFFF  P::::P     P:::::P  L:::::L
	#   F:::::F               P::::P     P:::::P  L:::::L
	#   F::::::FFFFFFFFFF     P::::PPPPPP:::::P   L:::::L
	#   F:::::::::::::::F     P:::::::::::::PP    L:::::L
	#   F:::::::::::::::F     P::::PPPPPPPPP      L:::::L
	#   F::::::FFFFFFFFFF     P::::P              L:::::L
	#   F:::::F               P::::P              L:::::L
	#   F:::::F               P::::P              L:::::L         LLLLLL
	# FF:::::::FF           PP::::::PP          LL:::::::LLLLLLLLL:::::L
	# F::::::::FF           P::::::::P          L::::::::::::::::::::::L
	# F::::::::FF           P::::::::P          L::::::::::::::::::::::L
	# FFFFFFFFFFF           PPPPPPPPPP          LLLLLLLLLLLLLLLLLLLLLLLL
	#
	#
	### Installation
	#
	# $$ make all
	#
	### Targets
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## Clean the repo
	@echo "🧹 Cleaning the repo..."
	rm -fr node_modules
	@echo "✔️  Done"

.PHONY: install
install: install-hooks install-go install-npm ## Lint everything

.PHONY: install-hooks
install-hooks:
	@echo "🎣 Installing git hooks"
	node scripts/install-hooks.js
	@echo "✔️  Done"

.PHONY: install-npm
install-npm: ## Install the local node dependencies
	@echo "📡 Installing local node dependencies..."
	npm ci
	@echo "✔️  Done"

.PHONY: install-go
install-go: ## Install the local go dependencies
	@echo "📡 Installing local go dependencies..."
	go install github.com/securego/gosec/cmd/gosec@master
	go install golang.org/x/lint/golint@master
	go get ./...
	@echo "✔️  Done"

.PHONY: vet
vet: ## Vet the code
	@echo "⚡ Vetting the code..."
	go vet -v ./...
	@echo "✔️  Done"

.PHONY: lint
lint: lint-go lint-markdown ## Lint everything

.PHONY: lint-go
lint-go: ## Lint the go code
	@echo "🔬 Linting the code..."
	golint -set_exit_status $(shell go list ./... | grep -v vendor)
	@echo "✔️  Done"

.PHONY: security
security: ## Inspect the code
	@echo "🔒 Checking code security..."
	gosec ./...
	@echo "✔️  Done"

.PHONY: build
build: ## Build the application
	@echo "🔨 Building the application..."
	go build .
	@echo "✔️  Done"

.PHONY: update-toc
update-toc: ## Update the README table of contents
	@echo "📑 Updating table of contents..."
	npm run toc
	@echo "✔️  Done"

.PHONY: lint-markdown
lint-markdown: ## Lint the markdown files
	@echo "🔬 Linting markdown files..."
	npm run lint-markdown
	@echo "✔️  Done"

.PHONY: spell-check
spell-check: ## Spellcheck markdown files
	@echo "📜 Spellchecking markdown files..."
	npm run spell-check
	@echo "✔️  Done"


.PHONY: all
all: clean install spell-check vet lint security build ## Run everything
