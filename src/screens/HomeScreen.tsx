import * as React from 'react';

import {
  Container,
  Paper,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  CircularProgress
} from '@material-ui/core';

import { gql } from 'apollo-boost';
import IEntry from '../../server/interfaces/Entry';
import { useQuery } from '@apollo/react-hooks';

import moment from 'moment';

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
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>Day</TableCell>
              <TableCell>Start</TableCell>
              <TableCell>End</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.map(entry => (
              <TableRow key={entry.id}>
                <TableCell>{entry.id}</TableCell>
                <TableCell>{moment(entry.day).format('DD/MM/YYYY')}</TableCell>
                <TableCell>{entry.start ? moment(entry.start).format('HH:mm') : '?'}</TableCell>
                <TableCell>{entry.end ? moment(entry.end).format('HH:mm') : '?'}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </Paper>
    </Container>
  );
};
