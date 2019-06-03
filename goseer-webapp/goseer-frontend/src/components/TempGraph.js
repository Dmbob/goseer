import React, { Component } from 'react';
import Websocket from "react-websocket";
import { BarChart, XAxis, YAxis, Bar } from "recharts";

const POLL_INTERVAL = 2000;

const styles = {
    container: {
        margin: "0 auto",
        width: "80%"
    }
}

class TempGraph extends Component {
    constructor(props) {
        super(props);

        this.state = {
            connected: false,
            tempData: []
        }
    }

    socketOpened() {
        console.log("Connected to server");
        this.setState({connected: true});
    }

    handleSocketMessage(data) {
        let jsonData = JSON.parse(data);
        this.setState({tempData: jsonData});
    }
    
    componentDidMount() { 
        setInterval(() => {
            let data = JSON.stringify({
                isRequest: true,
                requestType: "systemtemps"
            });

            this.refWebSocket.sendMessage(data);
        }, POLL_INTERVAL);
    }

    render() {
        return (
            <div style={styles.container}>
                {this.state.tempData.length > 0 && 
                    <div>
                        <BarChart width={700} height={600} data={this.state.tempData}>
                            <XAxis dataKey="adapter" />
                            <YAxis 
                                tickFormatter={(tick) => (`${tick} ºC`)}
                                domain={[0, parseInt(this.state.tempData[0].criticalTemp)]}
                            />
                            <Bar dataKey="currentTemp" fill="#8884d8" label/>
                        </BarChart>
                    </div>
                }
                <Websocket url="ws://192.168.1.2:8080/"
                    onOpen={this.socketOpened.bind(this)}
                    onMessage={this.handleSocketMessage.bind(this)}
                    ref={
                        Websocket => {
                            this.refWebSocket = Websocket;
                        }
                    }
                />
            </div>
        )
    }
}

export default TempGraph;