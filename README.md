# SCM - Go Microservices Boilerplate with GraphQL, MongoDB, and gRPC

SCM is a Go microservices boilerplate that provides a foundation for building scalable and modular applications using GraphQL, MongoDB, and gRPC. It aims to provide a structured and efficient starting point for developing microservices-based applications.

## Features

- Implements a microservices architecture with Go.
- Utilizes GraphQL for efficient and flexible API development.
- Integrates with MongoDB for data storage and retrieval.
- Supports gRPC for inter-service communication.
- Provides a Docker-based development environment for easy setup and deployment.

## Prerequisites

Make sure you have the following dependencies installed:

- [Go Kit](https://github.com/go-kit/kit)
- [Kit](https://github.com/kujtimiihoxha/kit)
- [GraphQL-Go](https://github.com/graphql-go/graphql)
- [mgo.v2](https://gopkg.in/mgo.v2)

## Getting Started

To run the SCM project, follow these steps:

1. Clone the repository.
2. Make sure you have Docker installed on your system.
3. Run the project using Docker Compose: `docker-compose up`.

## Usage

Once the project is up and running, you can interact with the services using the following commands:

- Call the database service: `curl -XPOST http://localhost:8800/connect -d '{"":""}'`.

## Contributing

Contributions are welcome! If you have any ideas, improvements, or bug fixes, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT). Please see the LICENSE file for more details.

## Acknowledgements

SCM is built upon the hard work and contributions of various open-source projects and libraries. We acknowledge and thank the authors and maintainers of those projects for their valuable work.

## Next Steps

- Implement additional microservices as per your project requirements.
- Customize the GraphQL schema and resolvers to suit your application needs.
- Extend functionality by integrating with other services or APIs.
- Add unit tests and ensure code coverage for reliable and robust applications.
- Optimize performance and scalability based on your specific use cases.
- Document the APIs and provide clear guidelines for developers and users.

Enjoy building your microservices-based application with SCM!
