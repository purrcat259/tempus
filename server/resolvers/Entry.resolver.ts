import { Resolver, Query } from 'type-graphql';
import Entry from '../models/Entry';

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
}
