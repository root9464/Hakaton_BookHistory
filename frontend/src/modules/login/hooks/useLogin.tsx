import { validateResult } from '@/shared/utils/utils';
import { useMutation } from '@tanstack/react-query';
import axios from 'axios';
import { z } from 'zod';
import { FormData } from '../components/LoginForm';

const _UserLoginSchema = z.object({
  message: z.string(),
  status: z.string(),
  name: z.string(),
  pasword: z.string(),
  role: z.string(),
});

type UserLoginResponse = z.infer<typeof _UserLoginSchema>;

export const useLogin = () =>
  useMutation({
    mutationKey: ['login'],
    mutationFn: async (userData: FormData) => {
      const { data, status, statusText } = await axios.get<UserLoginResponse>('/api/auth/authorize', {
        params: userData,
      });
      if (status !== 200) {
        throw new Error(`${status}: ${statusText}`);
      }
      return validateResult(data, _UserLoginSchema);
    },
  });
