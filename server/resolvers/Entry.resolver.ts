import { Resolver, Query, Mutation, Arg, Ctx } from 'type-graphql';
import Entry, { AddEntryInput } from '../models/Entry';
import { Context } from 'apollo-server-core';
import * as sequelize from 'sequelize';
import { padMonthNum } from '../util';

@Resolver(of => Entry)
export default class EntryResolver {
  @Query(() => Entry)
  public async entry(): Promise<Entry> {
    const entry = await Entry.findOne();
    if (!entry) {
      throw new Error('Entry not found');
    }
    return entry;
  }

  @Query(() => [Entry])
  public async allEntries(): Promise<Entry[]> {
    const entries = await Entry.findAll();
    return entries;
  }

  @Query(() => [Entry])
  public async entriesByMonth(
    @Arg('month') month: number // TODO: validate
  ): Promise<Entry[]> {
    const paddedMonth: string = padMonthNum(month);
    const entries = await Entry.findAll({
      // Issue:
      // https://github.com/sequelize/sequelize/issues/11241
      // @ts-ignore.
      where: sequelize.where(sequelize.fn('strftime', '%m', sequelize.col('day')), paddedMonth)
      // where: sequelize.where(sequelize.fn('month', sequelize.col('day')), month)
    });
    return entries;
  }

  // @Mutation()
  // public async insertEvent(@Arg('data') newEntryData: AddEntryInput, @Ctx() ctx: Context): Entry {}
}
