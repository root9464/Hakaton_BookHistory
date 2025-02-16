import { validateResult } from '@/shared/utils/utils';
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { z } from 'zod';

const UserSchema = z.object({
  userID: z.number(),
  cords: z.object({
    lat: z.number(),
    lon: z.number(),
  }),
});

const CoordinatesSchema = z.array(UserSchema);

type CoordinatesResponse = z.infer<typeof CoordinatesSchema>;

export const useCords = () =>
  useQuery({
    queryKey: ['coordinates'],
    queryFn: async () => {
      const { data, status, statusText } = await axios.get<CoordinatesResponse>('/api/exapi/get');
      if (status !== 200) {
        throw new Error(`${status}: ${statusText}`);
      }

      return validateResult(data, CoordinatesSchema);
    },

    staleTime: 1000 * 60,
  });
