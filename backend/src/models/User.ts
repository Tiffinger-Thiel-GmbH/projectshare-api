import { getModelForClass, prop, Ref } from '@typegoose/typegoose';

enum Gender {
  MALE = 'male',
  FEMALE = 'female'
}

class Address {
  @prop({ required: true }) city!: string;
  @prop() country?: string;
  @prop() zip?: string;
}

class PhoneNumber {
  @prop({ required: true }) phoneNumber!: number;
  @prop() name?: string;
}

export class User {
  @prop({ required: true }) firstName!: string;
  @prop({ required: true }) lastName!: string;
  @prop({ required: true, index: true, unique: true }) mail!: string;
  @prop({ default: Date.now }) date?: Date;
  @prop() active?: boolean;
  @prop({ enum: Gender }) gender?: Gender;
  @prop() address?: Address;
  @prop({ type: PhoneNumber }) phones?: PhoneNumber[];
  @prop({ ref: User }) friends?: Ref<User>[];
}

export const UserModel = getModelForClass(User);
