import { Sequelize } from 'sequelize-typescript';
import Entry from './models/Entry';

const sequelize = new Sequelize({
  database: 'tempus',
  dialect: 'sqlite',
  username: 'root',
  password: '',
  storage: ':memory:'
});

export default async () => {
  sequelize.addModels([Entry]);
  await sequelize.sync({ force: true });
  return sequelize;
};
