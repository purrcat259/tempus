import { Resolver, Query, Arg, Mutation } from 'type-graphql';
import Entry, { CompleteEntryInput } from '../models/Entry';
// import { padMonthNum } from '../util';
import { connection } from '../db';

@Resolver(() => Entry)
export default class EntryResolver {
  private entryRepository = connection.getRepository(Entry);

  @Query(() => Entry)
  public async entry(): Promise<Entry> {
    const entry = await this.entryRepository.findOne();
    if (!entry) {
      throw new Error('Entry not found');
    }
    return entry;
  }

  @Query(() => [Entry])
  public async allEntries(): Promise<Entry[]> {
    const entries = await this.entryRepository.find();
    return entries;
  }

  @Query(() => [Entry])
  public async entriesByMonth(
    @Arg('month') month: number // TODO: validate
  ): Promise<Entry[]> {
    // const paddedMonth: string = padMonthNum(month);
    // const entries = await this.entryRepository.find({
    //   // Issue:
    //   // https://github.com/sequelize/sequelize/issues/11241
    //   // @ts-ignore.
    //   where: sequelize.where(sequelize.fn('strftime', '%m', sequelize.col('start')), paddedMonth)
    //   // where: sequelize.where(sequelize.fn('month', sequelize.col('day')), month)
    // });

    // TODO: Implement month filter query
    const entries = await this.entryRepository.createQueryBuilder('entry').getMany();

    return entries;
  }

  @Mutation(() => Entry)
  public async completeEvent(@Arg('data') completeEntry: CompleteEntryInput): Promise<Entry> {
    const incompleteEntry: Entry | undefined = await this.entryRepository.findOne(completeEntry.id);
    if (incompleteEntry === undefined) {
      throw new Error(`Entry of ID: ${completeEntry.id} does not exist`);
    }
    incompleteEntry.end = completeEntry.end;
    await this.entryRepository.save(incompleteEntry);
    return incompleteEntry;
  }
}
