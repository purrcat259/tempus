import 'reflect-metadata';
import { ApolloServer } from 'apollo-server';
import * as dotenv from 'dotenv';
import { buildSchema } from 'type-graphql';

import initDB from './db';
import EntryResolver from './resolvers/Entry.resolver';
import Entry from './models/Entry';

dotenv.config();

interface Env {
  port: number;
  introspection: boolean;
  playground: boolean;
}

const env: Env = {
  port: (process.env.PORT ? parseInt(process.env.PORT) : null) || 4000,
  introspection: process.env.APOLLO_INTROSPECTION ? true : false,
  playground: process.env.APOLLO_PLAYGROUND ? true : false
};

const main = async () => {
  const connection = await initDB();

  const now = new Date();
  const later = new Date();
  const afterLunch = new Date();
  const headingHome = new Date();
  later.setHours(later.getHours() + 2);
  afterLunch.setHours(later.getHours() + 1);
  headingHome.setHours(afterLunch.getHours() + 4);
  const tomorrow = new Date();
  tomorrow.setDate(tomorrow.getDate() + 1);

  console.log(now, later, tomorrow);

  const testEntry = new Entry();
  testEntry.start = afterLunch;
  testEntry.end = headingHome;
  testEntry.type = 'work';

  await connection.manager.save(testEntry);

  const testEntry2 = new Entry();
  testEntry2.start = tomorrow;
  testEntry2.type = 'holiday';

  await connection.manager.save(testEntry2);

  const schema = await buildSchema({
    emitSchemaFile: {
      path: __dirname + '/schema.gql',
      commentDescriptions: true
    },
    resolvers: [EntryResolver]
  });

  const server: ApolloServer = new ApolloServer({
    introspection: env.introspection,
    playground: env.playground,
    schema
  });

  server.listen(env.port).then(({ url }) => console.log(`Server ON at: ${url}`));
};

(async () => {
  try {
    await main();
  } catch (e) {
    console.error(e);
  }
})();
