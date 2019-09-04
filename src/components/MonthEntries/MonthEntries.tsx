import * as React from 'react';
import { Typography, Box, CircularProgress } from '@material-ui/core';
import EntriesTable from '../EntriesTable/EntriesTable';

import { monthNameFromNumber } from '../../../server/util';
import { useQuery } from '@apollo/react-hooks';
import { gql } from 'apollo-boost';
import IEntry from '../../../server/interfaces/Entry';
import EntriesList from '../EntriesList/EntriesList';

interface IProps {
  month: number;
}

const entriesByMonthQuery = gql`
  query EntriesByMonth($month: Float!) {
    entriesByMonth(month: $month) {
      id
      start
      end
    }
  }
`;

export default (props: IProps) => {
  const { loading, error, data } = useQuery(entriesByMonthQuery, { variables: { month: props.month } });
  if (loading) {
    return <CircularProgress />;
  }
  if (error) {
    return <p>{JSON.stringify(error)}</p>;
  }

  const monthName: string = monthNameFromNumber(props.month);
  const rows: IEntry[] = data.entriesByMonth.map((entry: any) => entry as IEntry);
  console.log(rows);

  return (
    <Box>
      <Typography>{monthName}</Typography>
      {/* {rows.length > 0 ? <EntriesTable entries={rows} /> : <p>No Entries yet</p>} */}
      {rows.length > 0 ? <EntriesList entries={rows} /> : <p>No Entries yet</p>}
    </Box>
  );
};
