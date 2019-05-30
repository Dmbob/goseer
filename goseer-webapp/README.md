# GoSeer Monitor - Webapp #

This is the frontend of the application. It will format and display all of the data served by the endpoints connected.

It will also store authentication information as well as endpoint information, such as what endpoint servers you are able to connect to.

### Overview ###

This project was developed using [React](https://reactjs.org/) for the frontend and [Go](https://golang.org) as the backend.

The following is a list of dependancies required to build the webapp:
* [Axios](https://github.com/axios/axios) - REST framework for react used to make requests to the backend.
* [React-Bootstrap](https://github.com/react-bootstrap/react-bootstrap) - UI library to make things look nice :)
* [React-Websocket](https://github.com/mehmetkose/react-websocket) - Websocket library for react to facilitate the webhook connections to the endpoints
* [Recharts](https://github.com/recharts/recharts) - Graphing libary used to graph the data from the endpoint in an easy-to-read way

To buid the dependencies for the frontend, run:  
`yarn install && yarn build`