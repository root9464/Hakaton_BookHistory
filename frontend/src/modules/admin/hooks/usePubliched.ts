import { validateResult } from '@/shared/utils/utils';
import { useMutation } from '@tanstack/react-query';
import axios from 'axios';
import { z } from 'zod';

const PublishedShema = z.object({
  status: z.enum(['published', 'draft']),
});

type SuccessResponse = z.infer<typeof PublishedShema>;

export const usePubliched = () =>
  useMutation({
    mutationKey: ['published'],
    mutationFn: async (id: string) => {
      const { data, status, statusText } = await axios.put<SuccessResponse>(`/api/application/${id}`, {
        status: 'published',
      });
      if (status !== 200) {
        throw new Error(`${status}: ${statusText}`);
      }
      return validateResult(data, PublishedShema);
    },
  });
