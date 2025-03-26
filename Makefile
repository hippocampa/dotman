# Makefile

SRC := main.go
OUT_NAME := dotman
BUILD_DIR := bin
OUT_PATH := $(BUILD_DIR)/$(OUT_NAME)

run:
	@go run $(SRC)

build:
	@mkdir -p $(BUILD_DIR)
	@go build -o $(OUT_PATH) $(SRC)
	@echo "Built at $(OUT_PATH)"

clean:
	@rm -rf $(OUT_PATH)
	@rm -rf ~/.dotfiles
	@echo "Cleaned build and dotfiles"

