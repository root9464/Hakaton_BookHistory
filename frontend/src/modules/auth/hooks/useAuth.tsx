import { validateResult } from '@/shared/utils/utils';
import { useMutation } from '@tanstack/react-query';
import axios from 'axios';
import { z } from 'zod';
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
