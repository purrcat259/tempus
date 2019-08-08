import { ApolloServer } from 'apollo-server';

import resolvers from './resolvers';
import typeDefs from './schemas';

import dotenv from 'dotenv';

dotenv.config();

const port: number = (process.env.PORT ? parseInt(process.env.PORT) : null) || 4000;

const server: ApolloServer = new ApolloServer({ resolvers, typeDefs });

server.listen(port).then(({ url }) => console.log(`Server ON at: ${url}`));
