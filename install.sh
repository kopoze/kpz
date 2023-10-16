#!/bin/bash

# Step 1: Go to "/tmp/" folder
cd /tmp/

# Step 2: Install Go binaries version 1.19 if it doesn't exist
if ! command -v go &> /dev/null; then
    echo "Go 1.19 is not installed. Installing..."
    wget https://golang.org/dl/go1.19.linux-amd64.tar.gz
    tar -C /usr/local -xzf go1.19.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    echo "Go 1.19 installed successfully."
fi

# Step 3: Clone the "kpz" repository
git clone https://github.com/kopoze/kpz

# Step 4: Install project dependencies
cd kpz
go get -d -v ./...

# Step 5: Install the project with "go install"
go install

# Step 6: Remove "/tmp/kpz" folder from "/tmp"
cd /tmp/
rm -rf kpz

# Provide instructions on how to run the installed project
echo "Installation completed successfully."
echo "You can now run the 'kpz' project by executing 'kpz' from your command line."

