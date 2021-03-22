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

.PHONY: install
install: ## Install the local dependencies
	go install github.com/securego/gosec/cmd/gosec@master
	go install golang.org/x/lint/golint@master
	go get ./...

.PHONY: vet
vet: ## Vet the code
	go vet -v ./...

.PHONY: lint
lint: ## Lint the code
	golint -set_exit_status $(shell go list ./... | grep -v vendor)

.PHONY: security
security: ## Inspect the code
	gosec ./...

.PHONY: build
build: ## Build the application
	go build .
