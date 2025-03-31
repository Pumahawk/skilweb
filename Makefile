.PHONY: backend-clean backend-run backend-build

ARG=""

PROJECT_NAME="Skilweb"

# Backend vars
SOURCE_BACKEND_ROOT=source-backend
BACKEND_TARGET_DIR=target/backend
BACKEND_TARGET_EXEC=$(BACKEND_TARGET_DIR)/main
BACKEND_GFS=$(shell git ls-files source-backend)

backend-build | $(BACKEND_TARGET_EXEC): $(BACKEND_GFS)
	@echo Start build Go backend $(PROJECT_NAME)
	cd $(SOURCE_BACKEND_ROOT) && go build -o ../$(BACKEND_TARGET_EXEC) main/main.go

backend-run: $(BACKEND_TARGET_EXEC)
	@$(BACKEND_TARGET_EXEC) $(ARG)

backend-clean:
	@if [ -d $(BACKEND_TARGET_DIR) ]; then echo Delete target directory; rm -r $(BACKEND_TARGET_DIR); fi;
