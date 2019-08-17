import { Sequelize } from 'sequelize-typescript';
import Entry from './models/Entry';

const sequelize = new Sequelize({
  database: 'tempus',
  dialect: 'sqlite',
  username: 'root',
  password: '',
  storage: ':memory:'
});

export default () => {
  sequelize.addModels([Entry]);
  sequelize.sync({ force: true });
  return sequelize;
};
