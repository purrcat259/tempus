import * as React from 'react';
import { List, ListItem, Grid, Typography, Box, IconButton } from '@material-ui/core';

import IEntry from '../../../server/interfaces/Entry';
import moment from 'moment';

import AddIcon from '@material-ui/icons/Add';

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
    <Box>
      {Array.from(entriesByDate).map(([day, entries]) => (
        <Grid container>
          <Grid item xs={3}>
            <Typography>Day {day}</Typography>
          </Grid>
          <Grid item xs={3}>
            {getList(entries)}
          </Grid>
          <Grid item xs={2}>
            {getAddButton(day)}
          </Grid>
        </Grid>
      ))}
    </Box>
  );
};

const getList = (entries: IEntry[]) => {
  return (
    <List>
      {entries.map((entry: IEntry) => (
        <ListItem key={`entry-${entry.id}`}>
          {moment(entry.start).format('HH:mm')}
          {` -> `}
          {entry.end ? moment(entry.end).format('HH:mm') : getAddEndTimeButton(entry)}
        </ListItem>
      ))}
    </List>
  );
};

const getAddButton = (day: number) => {
  return (
    <IconButton aria-label="add">
      <AddIcon />
    </IconButton>
  );
};

const getAddEndTimeButton = (entry: IEntry) => {
  return (
    <IconButton aria-label="add">
      <AddIcon />
    </IconButton>
  );
};
