# SCM - Go Microservices Boilerplate with GQL MongoDB & gRPC

SCM is a Go microservices boilerplate that provides a foundation for building scalable and modular applications using GraphQL, MongoDB, and gRPC. It includes the necessary components and libraries to get started quickly.

## Prerequisites

Make sure you have the following prerequisites installed:

- [Go Kit](https://github.com/go-kit/kit): A toolkit for building microservices in Go.
- [GrantZheng/kit](https://github.com/GrantZheng/kit): A library that complements Go Kit and provides additional utilities.
- [GraphQL](https://github.com/graphql-go/graphql): An implementation of GraphQL in Go.
- [MongoDB](https://gopkg.in/mgo.v2): The official MongoDB driver for Go.

## Usage

1. Clone the repository: `git clone https://github.com/your-repo.git`
2. Install the project dependencies using [Glide](https://glide.sh/): `glide install`
3. Run the project: `docker-compose up`
4. Call the DB service using curl: `curl -XPOST http://localhost:8800/connect -d '{"":""}'`

## Resources

To learn more about the components used in this project, you can refer to the following resources:

- [Go Kit](https://github.com/go-kit/kit): Official documentation for Go Kit.
- [GrantZheng/kit](https://github.com/GrantZheng/kit): Documentation for GrantZheng's kit library.
- [GraphQL in Go](https://github.com/graphql-go/graphql): Official documentation for GraphQL in Go.
- [MongoDB Go Driver](https://gopkg.in/mgo.v2): Official documentation for the MongoDB Go driver.

## Project Importance

The SCM boilerplate provides a solid foundation for building microservices-based applications with a focus on scalability, modularity, and ease of development. It integrates popular technologies such as GraphQL, MongoDB, and gRPC, allowing developers to quickly build robust and efficient applications.

## What's Next

The project is actively maintained and open to contributions. Some potential areas of improvement and future enhancements include:

- Adding more examples and usage scenarios
- Enhancing security measures
- Integrating additional databases or message brokers
- Providing more extensive documentation and tutorials

Feel free to contribute to the project by submitting pull requests or opening issues for bug reports or feature requests.

