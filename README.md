## Go-Url-Shortener

This project demonstrates basic CRUD operations using the Go programming language.

## Contents

1. [Installation](#installation)
2. [Usage](#usage)
3. [Test](#test)
4. [Contributing](#contributing)
5. [License](#license)

## Installation

1. Clone this project: `git clone https://github.com/afurgapil/go-url-shortener.git`
2. Navigate to the project directory: `cd go-url-shortener`
3. Create the _urls_ table on your MySQL server by running the `db-schema.sql` file.
4. Install dependencies: `go mod tidy`
5. Check `.env` file

## Usage

1. Run the server: `go run cmd/go-url-shortener/main.go`
2. Access the API using an HTTP client.

## Test

Test files are located in the _test_ folder.

To run all tests `go test ./test/... -v`

## Contributing

If you encounter any issues or have suggestions for improvements, please feel free to contribute. Your feedback is highly appreciated and contributes to the learning experience.

> I especially need help in automating test processes. I have an automation experiment in .github/workflows/go.yaml, but I couldn't get the docker, mysql and test process to work properly. your help on this would make me very happy!

## License

This project is licensed under the [MIT License](LICENSE). For more information, please refer to the LICENSE file.
