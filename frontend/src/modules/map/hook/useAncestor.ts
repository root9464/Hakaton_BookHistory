import { DataSchema } from '@/modules/admin/hooks/useOrder';
import { SuccessResponse } from '@/shared/types/zod';
import { validateResult } from '@/shared/utils/utils';
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { z } from 'zod';

const AncestorShema = SuccessResponse.extend({
  data: DataSchema,
});

type AncestorResponse = z.infer<typeof AncestorShema>;

export const useAncestor = (id: string) =>
  useQuery({
    queryKey: ['ancestors', id],
    queryFn: async () => {
      const { data, status, statusText } = await axios.get<AncestorResponse>(`/api/application/${id}`);
      if (status !== 200) {
        throw new Error(`${status}: ${statusText}`);
      }

      return validateResult(data, AncestorShema);
    },
    select: (data) => data.data,
    staleTime: 1000 * 60,
  });
