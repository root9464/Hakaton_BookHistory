import { SuccessResponse } from '@/shared/types/zod';
import { validateResult } from '@/shared/utils/utils';
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { z } from 'zod';

const ImageSchema = z.object({
  id: z.string().optional(),
  name: z.string().optional(),
  owner_id: z.string().optional(),
  owner_type: z.string().optional(),
});

const RewardSchema = z.object({
  id: z.string(),
  name: z.string(),
  image: ImageSchema,
});

const FileSchema = z.object({
  id: z.string(),
  name: z.string(),
  owner_id: z.string(),
  owner_type: z.string(),
});

export const DataSchema = z.object({
  id: z.string(),
  sender_id: z.string(),
  fio: z.string(),
  geom: z.string(),
  date_of_birth: z.string(),
  place_of_birth: z.string(),
  name_of_millitary_commissariat: z.string(),
  millitary_rank: z.string(),
  date_of_death: z.string(),
  burial_place: z.string(),
  biographical_facts: z.string(),
  status: z.string(),
  files: z.array(FileSchema),
  rewards: z.array(RewardSchema),
});

const OrderShema = SuccessResponse.extend({
  data: z.array(DataSchema),
});

type GetOrderResponse = z.infer<typeof OrderShema>;

export const useOrder = () =>
  useQuery({
    queryKey: ['order'],
    queryFn: async () => {
      const { data, status, statusText } = await axios.get<GetOrderResponse>('/api/application/status/draft');
      if (status !== 200) {
        throw new Error(`${status}: ${statusText}`);
      }
      return validateResult(data, OrderShema);
    },
    staleTime: 1000 * 60,
    refetchOnWindowFocus: false,
    refetchOnMount: false,
  });
