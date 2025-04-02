.PHONY: clean be-build fe-install fe-build

ARG=""

PROJECT_NAME="Skilweb"

# Frontend vars
SOURCE_FE_ROOT=source-frontend
FE_TARGET_DIR=$(SOURCE_FE_ROOT)/dist
FE_FS=$(shell git ls-files source-frontend)

# Backend vars
SOURCE_BE_ROOT=source-backend
BE_TARGET_DIR=target/backend
BE_TARGET_EXEC=$(BE_TARGET_DIR)/main
BE_FS=$(shell git ls-files source-backend)

fe-build: $(SOURCE_FE_ROOT)/node_modules
	@cd $(SOURCE_FE_ROOT) && npm run build

fe-install:
	@cd $(SOURCE_FE_ROOT) && npm install --loglevel verbose

$(SOURCE_FE_ROOT)/node_modules: $(SOURCE_FE_ROOT)/package.json
	$(MAKE) fe-install

$(FE_TARGET_DIR) $(FE_TARGET_DIR)/index.html: $(FE_FS)
	$(MAKE) fe-build

be-build:
	@echo Start build Go bk $(PROJECT_NAME)
	cd $(SOURCE_BE_ROOT) && go build -o ../$(BE_TARGET_EXEC) main/main.go

$(BE_TARGET_EXEC): $(BE_FS)
	$(MAKE) be-build

clean:
	git clean -fdx
