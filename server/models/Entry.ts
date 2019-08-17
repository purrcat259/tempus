import { Table, Column, Model, PrimaryKey, DataType, AutoIncrement } from 'sequelize-typescript';
import { ObjectType, Field } from 'type-graphql';

@ObjectType()
@Table
export default class Entry extends Model<Entry> {
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
