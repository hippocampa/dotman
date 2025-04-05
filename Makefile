# Makefile

SRC := main.go
OUT_NAME := dotman
BUILD_DIR := bin
OUT_PATH := $(BUILD_DIR)/$(OUT_NAME)
GO_FILES := $(shell find . -name "*.go")

run:
	@go run $(SRC)

build:
	@gofmt -s -w $(GO_FILES)
	@mkdir -p $(BUILD_DIR)
	@go build -o $(OUT_PATH) $(SRC)
	@echo "Built at $(OUT_PATH)"

clean:
	@rm -rf $(OUT_PATH)
	@rm -rf ~/.dotfiles
	@echo "Cleaned build and dotfiles"

dev:
	@make clean
	@make build
	@touch ~/Desktop/lalalalostyou.md
	@touch ~/Desktop/testfile.md
	@echo "All my demons run wild\nAll my demons have your smile" >> ~/Desktop/lalalalostyou.md
	@./bin/dotman init
	@./bin/dotman track ~/Desktop/testfile.md
	@./bin/dotman track ~/Desktop/lalalalostyou.md

