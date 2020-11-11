import express from 'express';
import request from 'supertest';
import helloRouter from './hello';

describe('hello', () => {
  let app: express.Express;

  beforeEach(() => {
    app = express();
    app.use(helloRouter);
  });

  it('works', async () => {
    await request(app).get('/').expect('Content-Type', /text/).expect(200, 'Hello World!');
  });
});
