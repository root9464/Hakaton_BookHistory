import { AncestorResponse, AncestorShema } from '@/modules/map/hook/useAncestor';
import { validateResult } from '@/shared/utils/utils';
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { forkJoin, from, lastValueFrom, map, switchMap } from 'rxjs';
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

const OrderShema = z.object({
  data: z.array(DataSchema),
});

type GetOrderResponse = z.infer<typeof OrderShema>;

export const useCombinedOrders = () =>
  useQuery({
    queryKey: ['combinedOrders'],
    queryFn: async () => {
      const orders$ = from(axios.get<GetOrderResponse>('/api/application/status/draft'));
      const combined$ = orders$.pipe(
        switchMap((ordersResponse) => {
          const { data: ordersData } = validateResult(ordersResponse.data, OrderShema);
          const ancestorRequests = ordersData.map((order) => from(axios.get<AncestorResponse>(`/api/application/${order.id}`)));
          return forkJoin(ancestorRequests).pipe(
            map((ancestorResponses) => {
              const ancestorData = ancestorResponses.map((response) => validateResult(response.data, AncestorShema).data);
              return ordersData.map((order, index) => ({
                ...order,
                ancestorDetails: ancestorData[index],
              }));
            }),
          );
        }),
      );

      return lastValueFrom(combined$);
    },
    staleTime: 1000 * 60 * 5,
    refetchOnWindowFocus: false,
    refetchOnMount: false,
  });
