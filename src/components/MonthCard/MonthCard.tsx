import * as React from 'react';
import { Card, CardHeader, IconButton } from '@material-ui/core';

import { monthNameFromNumber } from '../../../server/util';
import ChevronRightIcon from '@material-ui/icons/ChevronRight';

interface IProps {
  month: number;
  onSetMonth: VoidFunction;
}

export default (props: IProps) => {
  const monthName: string = monthNameFromNumber(props.month);
  return (
    <Card>
      <CardHeader
        title={monthName}
        // subheader="TODO number of entries"
        action={
          <IconButton aria-label="expand" onClick={() => props.onSetMonth()}>
            <ChevronRightIcon />
          </IconButton>
        }
      />
    </Card>
  );
};
