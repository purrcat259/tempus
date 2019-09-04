import * as React from 'react';
import CompletedChart from './CompletedChart/CompletedChart';
import { Typography, Box } from '@material-ui/core';

export default () => {
  return (
    <Box>
      <Typography>Statistics</Typography>
      <CompletedChart />
    </Box>
  );
};
