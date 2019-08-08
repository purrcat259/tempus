import { ApolloServer } from 'apollo-server';

import resolvers from './resolvers';
import typeDefs from './schemas';

import dotenv from 'dotenv';

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

const server: ApolloServer = new ApolloServer({
  resolvers,
  typeDefs,
  introspection: env.introspection,
  playground: env.playground
});

server.listen(env.port).then(({ url }) => console.log(`Server ON at: ${url}`));
