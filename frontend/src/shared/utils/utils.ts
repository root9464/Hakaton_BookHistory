import { clsx, type ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';
import { z, ZodError } from 'zod';

export const validateResult = <T, U>(data: U, resType: z.ZodType<T>) => {
  try {
    return resType.parse(data);
  } catch (error) {
    console.error('Validation error:', (error as ZodError).toString());
    throw error;
  }
};

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}
