import React from 'react';
import TempGraph from "./components/TempGraph";
import Toolbar from "./components/Toolbar";

function App() {
	return (
		<div className="App">
			<Toolbar />
			<TempGraph />
		</div>
	);
}

export default App;
