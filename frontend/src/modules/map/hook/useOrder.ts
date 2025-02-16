import { SuccessResponse } from '@/shared/types/zod';
import { validateResult } from '@/shared/utils/utils';
import { useMutation } from '@tanstack/react-query';
import axios from 'axios';
import { z } from 'zod';

type CreateOrderResponse = z.infer<typeof SuccessResponse>;

export const useOrder = () =>
  useMutation({
    mutationKey: ['order'],
    mutationFn: async (orderData: FormData) => {
      const { data, status, statusText } = await axios.post<CreateOrderResponse>('/api/application', orderData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });

      if (status !== 200) {
        throw new Error(`${status}: ${statusText}`);
      }
      return validateResult(data, SuccessResponse);
    },
  });
