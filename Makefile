APP_NAME = jobtrack

PLATFORMS_LINUX = "linux amd64" "linux arm64" "linux arm"
PLATFORMS_MAC = "darwin amd64" "darwin arm64"
PLATFORMS_WINDOWS = "windows amd64" "windows arm64"

.PHONY: all
all: build-linux build-mac build-windows

.PHONY: build
build:
	@echo "Detecting platform...";
	@GOOS=$(shell go env GOOS); \
	GOARCH=$(shell go env GOARCH); \
	echo "Building for $$GOOS-$$GOARCH"; \
	GOOS=$$GOOS GOARCH=$$GOARCH go build -o "$(APP_NAME)" .; \
	sudo mv "$(APP_NAME)" /usr/local/bin/; \
	echo "$(APP_NAME) installed to /usr/local/bin/"

.PHONY: build-linux
build-linux:
	@echo "Building Linux binaries..."
	@for PLATFORM in $(PLATFORMS_LINUX); do \
		set -- $$PLATFORM; \
		GOOS=$$1 GOARCH=$$2; \
		OUTPUT_DIR="$(APP_NAME)-$$GOOS-$$GOARCH"; \
		mkdir "$$OUTPUT_DIR"; \
		cp README.md "$$OUTPUT_DIR"; \
		GOOS=$$1 GOARCH=$$2 go build -o "$$OUTPUT_DIR/$(APP_NAME)" .; \
		zip -r "$$OUTPUT_DIR.zip" "$$OUTPUT_DIR"; \
		rm -rf "$$OUTPUT_DIR"; \
	done

# Build and Zip macOS
.PHONY: build-mac
build-mac:
	@echo "Building macOS binaries..."
	@for PLATFORM in $(PLATFORMS_MAC); do \
		set -- $$PLATFORM; \
		GOOS=$$1 GOARCH=$$2; \
		OUTPUT_DIR="$(APP_NAME)-$$GOOS-$$GOARCH"; \
		mkdir "$$OUTPUT_DIR"; \
		cp README.md "$$OUTPUT_DIR"; \
		GOOS=$$1 GOARCH=$$2 go build -o "$$OUTPUT_DIR/$(APP_NAME)" .; \
		zip -r "$$OUTPUT_DIR.zip" "$$OUTPUT_DIR"; \
		rm -rf "$$OUTPUT_DIR"; \
	done

# Build and Zip Windows
.PHONY: build-windows
build-windows:
	@echo "Building Windows binaries..."
	@for PLATFORM in $(PLATFORMS_WINDOWS); do \
		set -- $$PLATFORM; \
		GOOS=$$1 GOARCH=$$2; \
		OUTPUT_DIR="$(APP_NAME)-$$GOOS-$$GOARCH"; \
		mkdir "$$OUTPUT_DIR"; \
		cp README.md "$$OUTPUT_DIR"; \
		GOOS=$$1 GOARCH=$$2 go build -o "$$OUTPUT_DIR/$(APP_NAME).exe" .; \
		zip -r "$$OUTPUT_DIR.zip" "$$OUTPUT_DIR"; \
		rm -rf "$$OUTPUT_DIR"; \
	done
