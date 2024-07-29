# Docker Container Monitor

## Description

This project is a command-line interface (CLI) that monitors running Docker containers and displays their details. It provides real-time information about the containers, such as their ID, name, status, and resource usage.

## Table of Contents

- [Example](#example)
- [Installation](#installation)
- [Usage](#usage)

## Example

![app](/docs/app.png)

## Installation

To install the project, follow these steps:

1. Clone the repository: `git clone https://github.com/buemura/container-monitor.git`
2. Navigate to the project directory: `cd container-monitor`
3. Install dependencies: `go mod tidy`

## Usage

To use the project, follow these steps:

1. Run the main script:

```bash
# Running go file
go run cmd/cli/main.go
```

or

```bash
# Running using makefile
make dev-cli
```

2. Configure the settings as needed.
3. The CLI will start monitoring running Docker containers and display their details in real-time.
