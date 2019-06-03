# GoSeer Monitor - Endpoint #

This is the endpoint server to be installed on the desired server. This creates a websocket server and serves information to the client.

### Overview ###

This project was developed using [Go 1.12.5](https://golang.org/dl/).

**Note**: GoSeer Endpoint requires `lm-sensors` to be installed on the server to work properly.

The following is a list of dependancies required to build GoSeer:
* [Gorilla/Websockets](https://github.com/gorilla/websocket) - Websocket framework for real-time connections.
* [Gosensors](https://github.com/ssimunic/gosensors) - Connector to connect `lm-sensors` to Go
* [Gorm](https://github.com/jinzhu/gorm) - ORM library for Go, used for database interactions

