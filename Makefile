# Define variables
APP_NAME = kpz
# VERSION = $(shell git describe --tags --abbrev=0 | sed 's/^v//')
VERSION = 0.1.0
BUILD_DIR = build
DEB_DIR = $(BUILD_DIR)/deb
INSTALL_DIR = $(DEB_DIR)/usr/local/bin
CONTROL_FILE = $(DEB_DIR)/DEBIAN/control
SERVICE_FILE = $(DEB_DIR)/etc/systemd/system/$(APP_NAME).service

# Define targets
all: build-deb

build-deb: build-binary prepare-service prepare-deb-files
	dpkg-deb --build $(DEB_DIR)
	dpkg-name $(BUILD_DIR)/deb.deb

build-binary:
	go build -o $(APP_NAME)

prepare-service:
	mkdir -p $(DEB_DIR)/etc/systemd/system/
	touch $(DEB_DIR)/etc/systemd/system/$(APP_NAME).service
	echo "[Unit]" > $(SERVICE_FILE)
	echo "Description=DevOps CLI toolkits made with Go" >> $(SERVICE_FILE)
	echo "" >> $(SERVICE_FILE)
	echo "[Service]" >> $(SERVICE_FILE)
	echo "ExecStart=/usr/local/bin/$(GO_BINARY)" >> $(SERVICE_FILE)
	echo "Restart=always" >> $(SERVICE_FILE)
	echo "User=root" >> $(SERVICE_FILE)
	echo "Group=root" >> $(SERVICE_FILE)
	echo "" >> $(SERVICE_FILE)
	echo "[Install]" >> $(SERVICE_FILE)
	echo "WantedBy=multi-user.target" >> $(SERVICE_FILE)

prepare-deb-files:
	mkdir -p $(DEB_DIR)/DEBIAN
	mkdir -p $(INSTALL_DIR)
	echo "Package: $(APP_NAME)" > $(CONTROL_FILE)
	echo "Version: $(VERSION)" >> $(CONTROL_FILE)
	echo "Architecture: amd64" >> $(CONTROL_FILE)
	echo "Maintainer: Hantsaniala Eléo <hantsaniala@gmail.com>" >> $(CONTROL_FILE)
	echo "Description: DevOps CLI toolkits made with Go" >> $(CONTROL_FILE)
	cp $(APP_NAME) $(INSTALL_DIR)

clean:
	rm -rf $(BUILD_DIR)
	rm -f $(APP_NAME)

.PHONY: all build-deb build-binary prepare-deb-files clean
