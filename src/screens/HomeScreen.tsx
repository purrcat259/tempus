import * as React from 'react';

import Button from '@material-ui/core/Button';
import { Container, Paper, Table, TableHead, TableRow, TableCell, TableBody } from '@material-ui/core';

export default () => {
  const entries: any[] = [{ id: 1, day: new Date(), start: new Date() }];

  return (
    <Container>
      <Paper style={{ padding: '2em' }}>
        <h1>Tempus</h1>
        <Button variant="contained" color="primary">
          Hello World
        </Button>
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
            {entries.map(entry => (
              <TableRow key={entry.id}>
                <TableCell>{entry.id}</TableCell>
                <TableCell>{entry.day.toISOString()}</TableCell>
                <TableCell>{entry.start.toISOString()}</TableCell>
                <TableCell>{entry.end ? entry.end.toISOString() : '?'}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </Paper>
    </Container>
  );
};
