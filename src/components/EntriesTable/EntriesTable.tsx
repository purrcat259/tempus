import * as React from 'react';
import { Table, TableHead, TableRow, TableCell, TableBody, List, ListItem } from '@material-ui/core';

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
          <TableCell>Day</TableCell>
          <TableCell>Start</TableCell>
          <TableCell>End</TableCell>
          <TableCell>Duration</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {Array.from(entriesByDate).map(([day, entries]) => (
          <TableRow>
            {/* entries.map((entry: IEntry) => (<TableCell rowSpan={entries.length}>{day}th</TableCell>
            <TableCell align="center" colSpan={3}>
              Test
            </TableCell>
            )) */}
            <TableCell rowSpan={entries.length}>DAY {day}</TableCell>
            <TableCell colSpan={3}>
              {entries.map((entry: IEntry) => (
                <TableRow>
                  <TableCell align="center">Test</TableCell>
                </TableRow>
              ))}
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};
