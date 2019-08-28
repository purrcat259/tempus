import * as React from 'react';
import * as ReactDOM from 'react-dom';
import HomeScreen from './screens/HomeScreen';

function App() {
  return <HomeScreen />;
}

ReactDOM.render(<App />, document.querySelector('#app'));
