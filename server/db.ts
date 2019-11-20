import Entry from './models/Entry';
import { createConnection, ConnectionOptions, Connection } from 'typeorm';

export let connection: Connection;

export default async () => {
  const options: ConnectionOptions = {
    type: 'sqlite',
    database: ':memory:',
    entities: [Entry],
    synchronize: true,
    logging: true
  };

  connection = await createConnection(options);

  return connection;
};
