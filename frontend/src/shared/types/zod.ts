import { z } from 'zod';

export const SuccessResponse = z.object({
  message: z.string(),
  status: z.string(),
});
