import { Table, Column, Model, PrimaryKey, DataType, AutoIncrement } from 'sequelize-typescript';
import { ObjectType, Field, InputType } from 'type-graphql';
import IEntry from '../interfaces/Entry';

@ObjectType()
@Table
export default class Entry extends Model<Entry> implements IEntry {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.INTEGER)
  public id!: number;

  @Field({ description: 'Day of work' })
  @Column
  day: Date;

  @Field({ description: 'Start Time', nullable: true })
  @Column
  start?: Date;

  @Field({ description: 'End Time', nullable: true })
  @Column
  end?: Date;
}

@InputType({ description: 'New Entry Data' })
export class AddEntryInput implements Partial<Entry> {
  @Field()
  day: Date;

  @Field({ nullable: true })
  start?: Date;

  @Field({ nullable: true })
  end?: Date;
}
