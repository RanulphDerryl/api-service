// types.ts
import { FastifyRequest, FastifyReply } from 'fastify';

export type Context = {
  req: FastifyRequest;
  reply: FastifyReply;
};

export type LoginResponse = {
  token: string;
  user: {
    id: number;
    email: string;
    name: string;
  };
};

export type User = {
  id: number;
  email: string;
  name: string;
};

export type CreateUserData = Omit<User, 'id'> & {
  password: string;
};

export type UpdateUserData = Partial<Omit<User, 'id'>> & {
  password?: string;
};