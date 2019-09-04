import * as React from 'react';
import { useState } from 'react';

import { Grid, Paper } from '@material-ui/core';

import MonthCard from '../components/MonthCard/MonthCard';
import MonthEntries from '../components/MonthEntries/MonthEntries';
import Statistics from '../components/Statistics/Statistics';

// TODO: Declare as constant
const monthNumbers: number[] = [];

for (let i = 1; i <= 12; i++) {
  monthNumbers.push(i);
}

export default () => {
  const [month, setMonth] = useState(1);
  return (
    <div style={{ flexGrow: 1 }}>
      <Grid container>
        <Grid item xs={2}>
          {monthNumbers.map((monthNum: number) => (
            <Grid item xs key={`month-grid-${monthNum}`}>
              <MonthCard month={monthNum} onSetMonth={() => setMonth(monthNum)} />
            </Grid>
          ))}
        </Grid>
        <Grid item xs={4}>
          <Paper
            style={{
              height: '100%',
              marginLeft: '1em',
              padding: '2em'
            }}
          >
            <MonthEntries month={month} />
          </Paper>
        </Grid>
        <Grid item xs={6}>
          <Paper
            style={{
              height: '100%',
              marginLeft: '1em',
              padding: '2em'
            }}
          >
            <Statistics />
          </Paper>
        </Grid>
      </Grid>
    </div>
  );
};
