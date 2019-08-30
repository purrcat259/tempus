import * as React from 'react';
import { Table, TableHead, TableRow, TableCell, TableBody } from '@material-ui/core';

import IEntry from '../../../server/interfaces/Entry';
import moment from 'moment';

interface IProps {
  entries: IEntry[];
}

export default (props: IProps) => {
  return (
    <Table>
      <TableHead>
        <TableRow>
          <TableCell>ID</TableCell>
          <TableCell>Day</TableCell>
          <TableCell>Start</TableCell>
          <TableCell>End</TableCell>
          <TableCell>Duration</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {props.entries.map(entry => (
          <TableRow key={entry.id}>
            <TableCell>{entry.id}</TableCell>
            <TableCell>{moment(entry.day).format('DD/MM/YYYY')}</TableCell>
            <TableCell>{entry.start ? moment(entry.start).format('HH:mm') : '?'}</TableCell>
            <TableCell>{entry.end ? moment(entry.end).format('HH:mm') : '?'}</TableCell>
            <TableCell>
              {entry.start && entry.end
                ? `${moment.duration(moment(entry.end).diff(moment(entry.start))).asHours()} Hours`
                : 'N/A'}
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};
