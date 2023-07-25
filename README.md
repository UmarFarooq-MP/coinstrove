# Coinscience Backend

Coinscience is a project that provides a backend implementation in Go for retrieving and comparing cryptocurrency prices across different exchanges. This backend allows viewers to see the price of a coin listed on various exchanges and find the exchange with the lowest price.

## Functionality

The backend code provides the following functionality:

- Retrieve the price of a specific coin from different exchanges.
- Compare the prices across exchanges and identify the exchange with the lowest price.
- Provide an API endpoint to retrieve the lowest price of a coin for a given exchange.

## Project Structure

The backend code follows the hexagonal architecture design pattern, also known as the ports and adapters architecture. The project structure is organized as follows:

- `cmd/`: Contains the entry point(s) for the application.
- `internal/`: Contains the core implementation of the application following the hexagonal architecture.
  - `app/`: Implements the application logic, use cases, and business rules.
  - `domain/`: Defines the domain models and interfaces.
  - `infrastructure/`: Implements the adapters and infrastructure code (e.g., external API clients, database access).
- `api/`: Contains the API handlers and routes for exposing the functionality to clients.
- `pkg/`: Contains reusable packages and utilities.
- `scripts/`: Contains scripts for building, testing, and other project-related tasks.

## Supported Exchanges

The backend currently supports the following cryptocurrency exchanges:

- Binance
- Gate.io
- Kraken
- Bitfinex
- Coinbase

## Technologies Used

The backend code is written in Go and utilizes the following technologies:

- Go programming language
- RESTful API development
- Third-party packages for interacting with the supported exchanges' APIs

## Usage

To use the backend code for the Coinscience project, follow these steps:

1. Clone the repository: `git clone https://github.com/UmarFarooq-MP/coinscience`
2. Install the necessary dependencies using a package manager like `go mod tidy`.
3. Run `docker-compose up -d` to up rabbitmttq in dockerize environment.
4. Build and run the application using `go run cmd/main.go`.

## API Endpoints
WORK IN PROGRESS

## Example Usage
WORK IN PROGRESS

## Contribution

Contributions to the Coinscience project are welcome! If you find any issues or have suggestions for improvements, feel free to submit a pull request or create an issue on the [GitHub repository](https://github.com/UmarFarooq-MP/coinscience).

## License

The Coinscience backend code is licensed under the [MIT License](https://opensource.org/licenses/MIT).
