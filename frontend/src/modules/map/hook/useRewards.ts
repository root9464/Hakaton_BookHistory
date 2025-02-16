import { SuccessResponse } from '@/shared/types/zod';
import { validateResult } from '@/shared/utils/utils';
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { z } from 'zod';

const ImageSchema = z.object({
  id: z.string(),
  name: z.string(),
  owner_id: z.string(),
  owner_type: z.string(),
});

const RewardSchema = z.object({
  id: z.string(),
  name: z.string(),
  image: ImageSchema,
});

const RewardsShema = SuccessResponse.extend({
  data: z.array(RewardSchema),
});

type GetRewardsResponse = z.infer<typeof RewardsShema>;

export const useRewards = () =>
  useQuery({
    queryKey: ['rewards'],
    queryFn: async () => {
      const { data, status, statusText } = await axios.get<GetRewardsResponse>('/api/reward');
      if (status !== 200) {
        throw new Error(`${status}: ${statusText}`);
      }

      return validateResult(data, RewardsShema);
    },

    staleTime: 1000 * 60,
  });
