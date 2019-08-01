import React from 'react';
import TempGraph from "./components/TempGraph";
import Toolbar from "./components/Toolbar";
import { Container, Row, Col } from 'react-bootstrap';

function App() {
	return (
		<div className="App">
			<Toolbar />
			<Container style={{maxWidth: "90%", marginTop: "30px"}}>
				<Row>
					<Col>
						<TempGraph size={{w: 500, h: 500}} />
					</Col>
				</Row>
			</Container>
			
		</div>
	);
}

export default App;
