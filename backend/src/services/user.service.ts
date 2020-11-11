import { UserDTO } from 'dto/user.dto';
import { User, UserModel } from 'models/User';

export async function findUsers(): Promise<User[]> {
  const users = await UserModel.find();
  return users;
}

export async function createUser(user: UserDTO): Promise<User> {
  const result = await UserModel.create({
    ...user,
    mail: `${user.firstName.toLowerCase()}.${user.lastName.toLowerCase()}@tiffinger-thiel.de`
  });
  return result;
}
