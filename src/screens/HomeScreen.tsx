import * as React from 'react';

import { Container, Paper, CircularProgress } from '@material-ui/core';

import { gql } from 'apollo-boost';
import IEntry from '../../server/interfaces/Entry';
import { useQuery } from '@apollo/react-hooks';

import EntriesTable from '../components/EntriesTable/EntriesTable';

const allEntriesQuery = gql`
  {
    allEntries {
      id
      day
      start
      end
    }
  }
`;

export default () => {
  const { loading, error, data } = useQuery(allEntriesQuery);
  if (loading) {
    return <CircularProgress />;
  }
  if (error) {
    return <p>{error}</p>;
  }
  const rows: IEntry[] = data.allEntries.map((entry: any) => entry as IEntry);

  return (
    <Container>
      <Paper style={{ padding: '2em' }}>
        <h1>Tempus</h1>
        <EntriesTable entries={rows} />
      </Paper>
    </Container>
  );
};
