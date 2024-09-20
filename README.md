# Hello Storage Proxy
[hello.app](https://hello.app/)'s storage service proxier used at [hello-desktop](https://github.com/hello-storage/hello-desktop) repo. 

Ensure you have installed:
- Docker
- Docker Compose
- Golang 1.22 or up (for host environment)
- Air (go install github.com/air-verse/air@latest)


Ensure that [backend-dev](https://github.com/hello-storage/hello-back) has already been started so that the network service and postgres container and volume exist.

You can set your _environment variables_ at .env file (change .env.example file name to .env)

You can set up your development environment as follows:

> Build and Start Services:

`$ make develop`

_Visit <local ip>:8181/api to check server status_

> View Logs:

`$ make logs`

> Stop Services:

`$ make stop-develop`

> Clean Up Services and Volumes:

`$ make down-develop`

> Run Tests:

`$ make test`

> Format Code:

`$ make fmt`

> Tidy Modules:

`$ make tidy`