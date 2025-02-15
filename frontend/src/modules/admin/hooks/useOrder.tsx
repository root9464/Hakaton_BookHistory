import { useMutation } from '@tanstack/react-query';
import { validateResult } from '@/shared/utils/utils';
import axios from 'axios';
import { z } from 'zod';
import { OrderFormData } from '../components/OrderForm';

const _UserOrderSchema = z.object({
  message: z.string(),
  status: z.string(),
  name: z.string(),
  files: z.array(z.instanceof(File)),
});

type UserOrderResponse = z.infer<typeof _UserOrderSchema>;

export const useOrder = () =>
  useMutation({
    mutationKey: ['order'],
    mutationFn: async (orderData: OrderFormData) => {
      const { data, status, statusText } = await axios.post<UserOrderResponse>('/api/auth/reward', orderData);
      if (status !== 200) {
        throw new Error(`${status}: ${statusText}`);
      }
      return validateResult(data, _UserOrderSchema);
    },
  });
