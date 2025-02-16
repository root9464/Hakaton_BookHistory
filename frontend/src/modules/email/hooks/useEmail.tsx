import { validateResult } from '@/shared/utils/utils';
import { useMutation } from '@tanstack/react-query';
import axios from 'axios';
import { z } from 'zod';
import { EmailData } from '@modules/email/components/EmailForm';

const _EmailSchema = z.object({
  message: z.string(),
  status: z.string(),
});

type EmailResponse = z.infer<typeof _EmailSchema>;

export const useEmail = () =>
  useMutation({
    mutationKey: ['email'],
    mutationFn: async (EmailData: EmailData) => {
      const { data, status, statusText } = await axios.post<EmailResponse>('http://127.0.0.1:6069/api/application/email', EmailData);
      if (status !== 200) {
        throw new Error(`${status}: ${statusText}`);
      }
      return validateResult(data, _EmailSchema);
    },
  });
