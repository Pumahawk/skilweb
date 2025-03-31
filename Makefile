.PHONY: bk-clean bk-run bk-build fr-install

ARG=""

PROJECT_NAME="Skilweb"

# Frontend vars
SOURCE_FR_ROOT=source-frontend
FBK_GFSR_TARGET_DIR=$(SOURCE_FR_ROOT)/dist
FR_FS=$(shell git ls-files source-frontend)

# Backend vars
SOURCE_BK_ROOT=source-backend
BK_TARGET_DIR=target/backend
BK_TARGET_EXEC=$(BK_TARGET_DIR)/main
BK_GFS=$(shell git ls-files source-frontend)

fr-install: $(SOURCE_FR_ROOT)/node_modules

$(SOURCE_FR_ROOT)/node_modules: 
	@cd $(SOURCE_FR_ROOT) && npm install --loglevel verbose

fr-build: fr-install $(FR_FS)
	@cd $(SOURCE_FR_RO
	OT) && npm run build

fr-run: fr-install
	@cd $(SOURCE_FR_ROOT) && npm run dev

fr-clean:
	@if [ -d $(FR_TARGET_DIR) ]; then echo Delete fr target directory; rm -r $(FR_TARGET_DIR); fi;

$(BK_TARGET_EXEC): $(BK_GFS)
	@echo Start build Go bk $(PROJECT_NAME)
	cd $(SOURCE_BK_ROOT) && go build -o ../$(BK_TARGET_EXEC) main/main.go

bk-build: $(BK_TARGET_EXEC)

bk-run: bk-build
	@$(BK_TARGET_EXEC) $(ARG)

bk-clean:
	@if [ -d $(BK_TARGET_DIR) ]; then echo Delete bk target directory; rm -r $(BK_TARGET_DIR); fi;
