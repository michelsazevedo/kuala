# Kuala

Job board micro-service to explore Hexagonal Architecture with Golang.

## Built With

- [Go](https://golang.org/)

Plus *some* of packages, a complete list of which is at [/master/go.mod](https://github.com/michelsazevedo/kuala/blob/master/go.mod).

## Instructions

### Dependencies

#### Running with Docker
[Docker](www.docker.com) is an open platform for developers and sysadmins to build, ship, and run distributed applications, whether on laptops, data center VMs, or the cloud.

If you haven't used Docker before, it would be good idea to read this article first: Install [Docker Engine](https://docs.docker.com/engine/installation/)

1. Install [Docker](https://www.docker.com/what-docker) and then [Docker Compose](https://docs.docker.com/compose/):

2. Run `docker compose build --no-cache` to build the images for the project.

3. Finally, run the local app with `docker-compose up web` and kuala will perform requests.

4. Aaaaand, you can run the automated tests suite running a `docker compose run --rm test` with no other parameters!

### References
[Ready for changes with Hexagonal Architecture](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749)

## License
Copyright © 2022
