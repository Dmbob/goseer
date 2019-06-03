# GoSeer Monitor #

This application consists of 2 major components: GoSeer Endpont and the GoSeer Webapp.

**GoSeer Endpoint is a binary that will run on Linux machines and provides an websocket endpoint to serve information over.**

**GoSeer Webapp is a web frontend/backend resposible for displaying information from GoSeer Endpoint as well as handling user authentication**

This project was developed using [React](https://reactjs.org/) (for the webapp) and [Go](https://golang.org/) (for the endpoint and webapp backend).

Please look in the directories of the individual application components for more information on dependencies and how to build them.

**This project is currently in early alpha stage and does not currently have all of the envisioned features implemented.**

## Building ##

### Webapp - Frontend ###
To build the React-based frontend, run these commands:

```
yarn install  
yarn build 
```

You will then have a `./build` directory. Copy the files from here and put them on the document root of your webserver

### Webapp - API ###
**Note**: You may want to install some sort of multiplexer like `tmux` to run the API so that you can leave the program running and exit the terminal.

To build the backend, go to the directory `./goseer-backend/cmd/backend-api` and run this command:
```
go build
```

Then rename `config/config.example.json` to `config.json` and fill the values in that file with your desired settings.

You can then run the `./backend-api` file and the API will be up and running.

### Endpoint ###
**Note**: You may want to install some sort of multiplexer like `tmux` to run the endpoint so that you can leave the program running and exit the terminal.

To build the backend, go to the directory `./cmd/endpoint` and run this command:
```
go build
```

Then all you need to do is run the `./endpoint` file and the endpoint will be running.