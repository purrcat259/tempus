import { Entity, PrimaryGeneratedColumn, Column } from 'typeorm';
import { ObjectType, Field, InputType } from 'type-graphql';
import { IEntry } from '../interfaces';

@ObjectType()
@Entity()
export default class Entry implements IEntry {
  @Field({ description: 'Entry Unique ID' })
  @PrimaryGeneratedColumn()
  public id!: number;

  @Field({ description: 'Type of Entry' })
  @Column()
  type: string;

  @Field({ description: 'Start Time', nullable: true })
  @Column()
  start: Date;

  @Field({ description: 'End Time', nullable: true })
  @Column({ nullable: true })
  end?: Date;
}

@InputType({ description: 'New Entry Data' })
export class AddEntryInput implements Partial<Entry> {
  @Field()
  start: Date;

  @Field()
  type: string;
}

@InputType({ description: 'Complete an existing Entry Data' })
export class CompleteEntryInput implements Partial<Entry> {
  @Field()
  id: number;

  @Field()
  end: Date;
}
