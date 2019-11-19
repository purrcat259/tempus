import Entry from './models/Entry';
import { createConnection, ConnectionOptions } from 'typeorm';

export default async () => {
  const options: ConnectionOptions = {
    type: 'sqlite',
    database: ':memory:',
    entities: [],
    logging: true
  };

  const connection = createConnection(options);

  return connection;
};
