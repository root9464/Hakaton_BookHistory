import { validateResult } from '@/shared/utils/utils';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import axios from 'axios';
import { z } from 'zod';
import { LoginFormData } from '../components/LoginForm';
import { FormData } from '../components/RegisterForm';

const _UserRegisterSchema = z.object({
  message: z.string(),
  status: z.string(),
});

type UserRegisterResponse = z.infer<typeof _UserRegisterSchema>;

export const useRegister = () =>
  useMutation({
    mutationKey: ['register'],
    mutationFn: async (userData: FormData) => {
      const { data, status, statusText } = await axios.post<UserRegisterResponse>('/api/auth/register', userData);
      if (status !== 200) {
        throw new Error(`${status}: ${statusText}`);
      }
      return validateResult(data, _UserRegisterSchema);
    },
  });

const TokenSchema = z.object({
  accessToken: z.string(),
  refreshToken: z.string(),
});

const UserSchema = z.object({
  id: z.string().uuid(),
  email: z.string().email(),
  password: z.string(),
  name: z.string(),
  surname: z.string(),
  patronymic: z.string().optional(),
  phone: z.string().regex(/^\+?\d{10,15}$/, 'Invalid phone number'),
  role: z.enum(['user', 'admin']),
});

const _UserLoginSchema = z.object({
  data: UserSchema,
  message: z.string(),
  status: z.string(),
  token: TokenSchema,
});

export type UserLoginResponse = z.infer<typeof _UserLoginSchema>;

export const useLogin = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: ['login'],
    mutationFn: async (userData: LoginFormData) => {
      const { data, status, statusText } = await axios.post<UserLoginResponse>('/api/auth/authorize', userData);
      if (status !== 200) {
        throw new Error(`${status}: ${statusText}`);
      }
      return validateResult(data, _UserLoginSchema);
    },

    onSuccess: (data) => {
      queryClient.setQueryData(['user'], data);
    },
  });
};
