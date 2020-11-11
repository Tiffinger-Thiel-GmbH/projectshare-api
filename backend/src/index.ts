import cors from 'cors';
import express from 'express';
import morgan from 'morgan';
import helloRouter from 'routes/hello';
import usersRouter from 'routes/users';
import 'utils/connection';

const app = express();
app.use(cors());
app.use(express.json());
app.use(morgan('tiny'));
const port = process.env.PORT || 3001;

app.use('/hello', helloRouter);
app.use('/users', usersRouter);

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});
