## Service for PDF files generation.

### Stack:

* [Fx](https://github.com/uber-go/fx)
* [Gin](https://github.com/gin-gonic/gin)
* [Swaggo](https://github.com/swaggo)
* [Testify](https://github.com/stretchr/testify)

### Commands:

#### Generate OpenAPI spec:

````shell
swag init --parseDependency
````

#### Docker build

````shell
docker build -t pdf-generator:latest .
````

#### Docker run

````shell
docker run -d -t -i -p 8080:8080 --name pdf-generator pdf-generator:latest
````

### Generated file example: