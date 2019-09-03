import { ApolloServer } from 'apollo-server';

import 'reflect-metadata';

import dotenv from 'dotenv';
// import db from './db';
import { buildSchema } from 'type-graphql';
import EntryResolver from './resolvers/Entry.resolver';
import { Sequelize } from 'sequelize-typescript';
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

// const GQLPATH = '/graphql';

const main = async () => {
  const sequelize = new Sequelize({
    database: 'tempus',
    dialect: 'sqlite',
    username: 'root',
    password: '',
    storage: ':memory:'
  });

  sequelize.addModels([Entry]);
  await sequelize.sync({ force: true });

  const now = new Date();
  const later = new Date();
  later.setHours(later.getHours() + 2);
  const tomorrow = new Date();
  tomorrow.setDate(tomorrow.getDate() + 1);

  console.log(now, later, tomorrow);

  let testEntry = new Entry({ day: now, start: now, end: later });
  await testEntry.save();
  testEntry = new Entry({ day: tomorrow });
  await testEntry.save();

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

main().catch(err => {
  console.error(err);
});
