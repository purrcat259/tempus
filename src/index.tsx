import * as React from 'react';
import * as ReactDOM from 'react-dom';
import HomeScreen from './screens/HomeScreen';

import ApolloClient from 'apollo-boost';
import { ApolloProvider } from '@apollo/react-hooks';

const client = new ApolloClient({
  uri: 'http://localhost:4444/gql'
});

function App() {
  return (
    <ApolloProvider client={client}>
      <HomeScreen />
    </ApolloProvider>
  );
}

ReactDOM.render(<App />, document.querySelector('#app'));
