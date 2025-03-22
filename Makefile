APP_NAME = jobtrack

PLATFORMS_LINUX = "linux amd64" "linux arm64" "linux arm"
PLATFORMS_MAC = "darwin amd64" "darwin arm64"
PLATFORMS_WINDOWS = "windows amd64" "windows arm64"

.PHONY: release
release: build-linux build-mac build-windows

.PHONY: build
build:
	@echo "Detecting platform...";
	@GOOS=$(shell go env GOOS); \
	GOARCH=$(shell go env GOARCH); \
	echo "Building for $$GOOS-$$GOARCH"; \
	GOOS=$$GOOS GOARCH=$$GOARCH go build -o "$(APP_NAME)" .;


.PHONY: install
install: build
	@echo "Installing $(APP_NAME)";
	@sudo mv "$(APP_NAME)" /usr/local/bin/;
	@echo "$(APP_NAME) installed to /usr/local/bin/";
	@echo "Installing man pages.....";
	@sudo cp "./man/jobtrack.1" /usr/share/man/man1/;
	@sudo cp "./man/jobtrack-create.1" /usr/share/man/man1/;
	@sudo cp "./man/jobtrack-list.1" /usr/share/man/man1/;
	@sudo cp "./man/jobtrack-delete.1" /usr/share/man/man1/;
	@sudo cp "./man/jobtrack-update.1" /usr/share/man/man1/;
	@sudo cp "./man/jobtrack-import.1" /usr/share/man/man1/;
	@sudo cp "./man/jobtrack-export.1" /usr/share/man/man1/;
	@echo "Install complete";

.PHONY: uninstall
uninstall:
	@echo "Uninstalling $(APP_NAME)...";
	@sudo rm -rf /usr/local/bin/$(APP_NAME);
	@rm -rf "$HOME/.local/share/jobtrack";
	@sudo rm -rf "/usr/share/man/man1/jobtrack.1";
	@sudo rm -rf "/usr/share/man/man1/jobtrack-create.1";
	@sudo rm -rf "/usr/share/man/man1/jobtrack-list.1";
	@sudo rm -rf "/usr/share/man/man1/jobtrack-delete.1";
	@sudo rm -rf "/usr/share/man/man1/jobtrack-update.1";
	@sudo rm -rf "/usr/share/man/man1/jobtrack-import.1";
	@sudo rm -rf "/usr/share/man/man1/jobtrack-export.1";
	@echo "Uninstall complete";

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
