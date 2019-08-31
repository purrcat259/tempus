import moment from 'moment';

export const padMonthNum = (month: number): string => `${month < 10 ? '0' : ''}${month}`;

export const monthNameFromNumber = (month: number) => {
  const paddedMonth: string = padMonthNum(month);
  const monthName: string = moment(paddedMonth, 'MM').format('MMMM');
  return monthName;
};
