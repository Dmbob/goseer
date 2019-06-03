import React, { Component } from "react";
import { Navbar, Button, Nav, NavDropdown } from "react-bootstrap";
import Form from "react-bootstrap/FormControl";

const styles = {
    toolbar: {
        paddingLeft: "5%",
        paddingRight: "5%",
    }
}

class Toolbar extends Component {
    render() {
        return (
            <Navbar style={styles.toolbar} expand="lg" bg="dark" variant="dark">
                <Navbar.Brand href="#home">GoSeer Monitor</Navbar.Brand>
                <Navbar.Toggle />
                <Navbar.Collapse className="justify-content-end">
                    <Navbar.Text>Logout</Navbar.Text>
                </Navbar.Collapse>
            </Navbar>
        )
    }
}

export default Toolbar;