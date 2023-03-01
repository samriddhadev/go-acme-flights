## A Reference Go Microservice implementation
The project provides a reference Microservice implementation in Go laguage.
This implementation tried to achieve a close resemble with the spring boot Microservice implementation (https://github.com/blueperf/acmeair-mainservice-springboot)

### Frameworks 
- RESTFul Interfaces: Gin Web Framework
- Dependency Injection: Google Wire
- PosgreSQL ORM: Bun

### To execute
~~~
go build -o /flightapp.exe  ./cmd/...
./flightapp.exe
~~~