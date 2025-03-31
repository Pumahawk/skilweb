.PHONY: run clean

PROJECT_NAME="Skilweb"

TARGET_DIR=target
TARGET_EXEC=$(TARGET_DIR)/main
GFS=$(shell git ls-files)

build | $(TARGET_EXEC): $(GFS)
	@echo Start build $(PROJECT_NAME)
	go build -o $(TARGET_EXEC) main/main.go

run: $(TARGET_EXEC)
	@$(TARGET_EXEC)

clean:
	@if [ -d $(TARGET_DIR) ]; then echo Delete target directory; rm -r $(TARGET_DIR); fi;
