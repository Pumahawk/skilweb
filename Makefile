.PHONY: be-clean be-run be-build fe-install

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
BE_GFS=$(shell git ls-files source-backend)

fe-install: $(SOURCE_FE_ROOT)/node_modules

$(SOURCE_FE_ROOT)/node_modules: 
	@cd $(SOURCE_FE_ROOT) && npm install --loglevel verbose

fe-build: fe-install $(FE_FS)
	@cd $(SOURCE_FE_ROOT) && npm run build

fe-run: fe-install
	@cd $(SOURCE_FE_ROOT) && npm run dev

fe-clean:
	@if [ -d $(FE_TARGET_DIR) ]; then echo Delete fr target directory; rm -r $(FE_TARGET_DIR); fi;

$(BE_TARGET_EXEC): $(BE_GFS)
	@echo Start build Go bk $(PROJECT_NAME)
	cd $(SOURCE_BE_ROOT) && go build -o ../$(BE_TARGET_EXEC) main/main.go

be-build: $(BE_TARGET_EXEC)

be-run: be-build
	@$(BE_TARGET_EXEC) $(ARG)

be-clean:
	@if [ -d $(BE_TARGET_DIR) ]; then echo Delete bk target directory; rm -r $(BE_TARGET_DIR); fi;
