import * as React from 'react';
import { Table, TableHead, TableRow, TableCell, TableBody } from '@material-ui/core';

import IEntry from '../../../server/interfaces/Entry';
import moment from 'moment';

interface IProps {
  entries: IEntry[];
}

export default (props: IProps) => {
  // Collate entries by day
  const entriesByDate: Map<number, IEntry[]> = new Map<number, IEntry[]>();
  props.entries.forEach((entry: IEntry) => {
    const date: number = new Date(entry.start).getDate();
    entriesByDate.get(date) === undefined
      ? entriesByDate.set(date, [entry])
      : (entriesByDate.get(date) as IEntry[]).push(entry);
  });
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
        {Array.from(entriesByDate).map(([day, entries]) => entries.map((entry: IEntry) => getRow({ entry })))}
      </TableBody>
    </Table>
  );
};

const getRow = (props: { entry: IEntry }) => {
  const entry = props.entry;
  return (
    <TableRow key={entry.id}>
      <TableCell>{entry.id}</TableCell>
      <TableCell>{moment(entry.start).format('Do MMMM')}</TableCell>
      <TableCell>{entry.start ? moment(entry.start).format('HH:mm') : '?'}</TableCell>
      <TableCell>{entry.end ? moment(entry.end).format('HH:mm') : '?'}</TableCell>
      <TableCell>
        {entry.start && entry.end
          ? `${moment.duration(moment(entry.end).diff(moment(entry.start))).asHours()} Hours`
          : 'N/A'}
      </TableCell>
    </TableRow>
  );
};
