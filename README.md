# GoSeer Monitor #

This application was created as a quick and simple way to montior temperatures and processes for Linux-based servers.

### Overview ###

This project was developed using [Go 1.12.5](https://golang.org/dl/).

**Note**: GoSeer requires `lm-sensors` to be installed on the server to work properly.

The following is a list of dependancies required to build GoSeer:
* [Gin](https://github.com/gin-gonic/gin) - Backend web framework used for the API
* [Gosensors](https://github.com/ssimunic/gosensors) - Connector to connect `lm-sensors` to Go
* [Gorm](https://github.com/jinzhu/gorm) - ORM library for Go, used for database interactions

