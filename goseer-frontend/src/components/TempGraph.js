import React, { Component } from 'react';
import axios from 'axios';
import { VictoryBar, VictoryChart, VictoryAxis } from "victory";

//axios.defaults.withCredentials = true;
//axios.defaults.auth = {username: "bob", password: "test"}

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
            tempData: [
                {
                    adapter: "Core 0",
                    currentTemp: 30,
                    highTemp: 80,
                    criticalTemp: 100
                }   
            ]
        }
    }
    
    componentDidMount() {
        setInterval(() => {
            axios.get("http://192.168.1.2:8080/api/v1/temperatures")
            .then(response => {
                let temps = response.data.temps;
                if(response.data.temps) {
                    temps.map((temp, index) => {
                        temps[index]["currentTemp"] = parseFloat(temps[index]["currentTemp"]);
                    });
                }
                
                
                this.setState({tempData: temps});
            }).catch(error => {
                console.log(error);
            });
        }, 2000);
    }

    render() {
        return (
            <div style={styles.container}>
                <h3>Current Temp: {parseFloat(this.state.tempData[0].currentTemp)+"ºC"}</h3>
                <VictoryChart domain={{y:[0, 30]}} domainPadding={100} height={200}>
                    <VictoryBar data={this.state.tempData} animate={{duration: 1000, onLoad: {duration: 1000}}} x="adapter" y="currentTemp" />
                </VictoryChart>
            </div>
        )
    }
}

export default TempGraph;