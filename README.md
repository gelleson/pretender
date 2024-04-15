# Pretender

Pretender is a versatile tool designed to generate fake server responses, ideal for testing and development purposes. It allows developers to simulate server behavior and responses without the need for an actual server, making it a valuable asset in a microservices architecture or when working with APIs.

## Features

- **Flexible Content Types**: Supports various content types including JSON, HTML, and plain text.
- **Customizable Responses**: Define default content to return for all requests.
- **Port Configuration**: Easily configure the port on which Pretender listens.
- **Pipe Support**: Read input data from a pipe, allowing dynamic response generation.

## Getting Started

### Prerequisites

Ensure you have the following installed:
- Go (version 1.22 or higher)

### Installation

To get Pretender up and running, follow these steps:

1. Clone the repository:
   ```sh
   git clone https://github.com/gelleson/pretender.git
   ```

2. Navigate to the project directory:
   ```sh
   cd pretender
   ```

3. Build the project:
   ```sh
   go build -o pretender
   ```

### Usage

Run Pretender with the following command:

```sh
./pretender --port 8080 --content-type "application/json" --default-content '{"message": "Hello, World!"}'
```

You can also pipe data directly into Pretender:

```sh
echo '{"message": "Hello, Pipe!"}' | ./pretender
```

### Configuration Flags

- `--port`, `-p`: The port to listen on (default: `8080`).
- `--content-type`, `-c`: The content type to use (default: `application/json`).
- `--default-content`, `-d`: The default content to return if no data is piped (default: "").
- `--prefork`, `-f`: Prefork the server (default: `false`).

## Development

To contribute to Pretender, you can follow the standard Git workflow:

1. Fork the repository.
2. Create a new feature branch.
3. Commit your changes.
4. Push to the branch.
5. Open a pull request.

### Running Tests

To run tests, execute the following command in the project directory:

```sh
go test ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

- Thanks to all the contributors who have helped shape Pretender into what it is today.
- Special thanks to the Go community for providing the tools and libraries that make projects like this possible.

---

For more information on Pretender, please visit the [official documentation](#) or the [GitHub repository](https://github.com/gelleson/pretender).
