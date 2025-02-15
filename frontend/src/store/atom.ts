import { atom } from 'jotai';
import { z } from 'zod';


export const UserSchema = z.object({
  id: z.string().uuid(),
  name: z.string(),
  surname: z.string().optional(),
  patronymic: z.string().optional(),
  email: z.string().email().optional(),
  password: z.string().min(8),
  phone: z.string().min(10).max(10).optional(),
});

export const ApplicationSchema = z.object({
  id: z.string().uuid(),
  userId: z.string().uuid(),
  status: z.enum(['pending', 'approved']),
  description: z.string().min(1),
});

// Типы
export type User = z.infer<typeof UserSchema>;
export type Application = z.infer<typeof ApplicationSchema>;

// Атомы для хранения данных
export const usersAtom = atom<User[]>([]);
export const applicationsAtom = atom<Application[]>([]);