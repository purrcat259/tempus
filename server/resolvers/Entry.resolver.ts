import { Resolver, Query, Mutation, Arg, Ctx } from 'type-graphql';
import Entry, { AddEntryInput } from '../models/Entry';
import { Context } from 'apollo-server-core';

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

  @Mutation()
  public async insertEvent(@Arg('data') newEntryData: AddEntryInput, @Ctx() ctx: Context): Entry {}
}
