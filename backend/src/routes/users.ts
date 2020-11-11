import express from 'express';
import { UserDTO } from 'dto/user.dto';
import { createUser, findUsers } from 'services/user.service';

const router = express.Router();

router.get('/', async (req, res) => {
  const users = await findUsers();
  res.send(users);
});

router.post('/', async (req, res) => {
  const user: UserDTO = req.body;
  const result = await createUser(user);
  res.send(result);
});

export default router;
