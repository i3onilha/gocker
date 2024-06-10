
# Gocker

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Gocker is a start up, minimal containerized Docker environment. It allows you to create and manage a new container using the same fundamental concepts as Docker. This project is perfect for those who want to get start a new Golang project with all tools necessary. 

## Features

- Container lifecycle management (create, start, stop, delete)
- Basic image management
- Networking capabilities
- Resource isolation using cgroups and namespaces
- CLI interface for container management

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Installation

### Prerequisites

- Go 1.16 or higher
- Linux operating system (required for namespaces and cgroups)

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/i3onilha/gocker.git
   cd gocker
   ```

2. Start new project:
   ```bash
   make dev 
   ```
## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new feature branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Create a new Pull Request.

Make sure to follow the [contribution guidelines](CONTRIBUTING.md) and maintain code quality with [Go Report Card](https://goreportcard.com/report/github.com/i3onilha/gocker).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Special thanks to the Go community for their contributions and support.
- Inspired by the principles and architecture of Docker.
