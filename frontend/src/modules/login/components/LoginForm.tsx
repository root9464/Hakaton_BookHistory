import { Button, Input } from '@heroui/react';
import { zodResolver } from '@hookform/resolvers/zod';
import { Link } from '@tanstack/react-router';
import { useForm } from 'react-hook-form';
import * as z from 'zod';
import { useLogin } from '../hooks/useLogin';

const FormShema = z.object({
  email: z.string().email('Некорректный email'),
  password: z
    .string()
    .min(8, 'Пароль должен содержать минимум 8 символов')
    .regex(/[A-Z]/, 'Пароль должен содержать хотя бы одну заглавную букву')
    .regex(/[0-9]/, 'Пароль должен содержать хотя бы одну цифру'),
  phone: z
    .string()
    .min(10, 'Номер телефона должен содержать 10 цифр')
    .max(10, 'Номер телефона должен содержать 10 цифр')
    .regex(/^\d{10}$/, 'Номер телефона должен содержать только цифры')
    .transform((val) => `+7${val}`),
  name: z.string(),
  role: z.string(),
});

export type FormData = z.infer<typeof FormShema>;
export const LoginForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: zodResolver(FormShema),
    defaultValues: {
      email: '',
      password: '',
      phone: '',
    },
  });

  const { data, mutate } = useLogin();
  console.log(data?.message);

  const onSubmit = (data: FormData) => mutate(data);

  return (
    <form onSubmit={handleSubmit(onSubmit)} className='flex flex-col'>
      <div className='flex flex-col gap-5'>
        <Input
          {...register('email')}
          label='Email'
          placeholder='Введите email'
          className='h-12 w-[315px]'
          errorMessage={errors.email?.message}
          isInvalid={!!errors.email}
        />
        <Input
          {...register('password')}
          label='Password'
          placeholder='Придумайте пароль'
          className='h-12 w-[315px]'
          type='password'
          errorMessage={errors.password?.message}
          isInvalid={!!errors.password}
        />
        <Input
          {...register('phone')}
          label='Phone'
          placeholder='Номер телефона'
          className='h-12 w-[315px]'
          errorMessage={errors.phone?.message}
          isInvalid={!!errors.phone}
        />
      </div>
      <div className='flex flex-col items-center justify-center place-self-center'>
        <div className='mt-[37px] flex flex-row gap-4 text-xs font-medium text-blue-600'>
          <Link to={'/login'}>Уже есть аккаунт</Link>
          <Link to='.'>Войти через гос услуги</Link>
        </div>
        <Button type='submit' className='mt-4 bg-uiDeepGray p-2 text-white'>
          Войти
        </Button>
      </div>
    </form>
  );
};
