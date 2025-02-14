import { Button, Input } from '@heroui/react';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import * as z from 'zod';

const schema = z.object({
  email: z.string().email('Некорректный email'),
  password: z
    .string()
    .min(8, 'Пароль должен содержать минимум 8 символов')
    .regex(/[A-Z]/, 'Пароль должен содержать хотя бы одну заглавную букву')
    .regex(/[0-9]/, 'Пароль должен содержать хотя бы одну цифру'),
  name: z.string().min(1, 'Имя обязательно для заполнения'),
  surname: z.string().min(1, 'Фамилия обязательна для заполнения'),
  patronymic: z.string().optional(),
  phone: z.string().regex(/^\+7\d{10}$/, 'Номер телефона должен быть в формате +7XXXXXXXXXX'),
});

type FormData = z.infer<typeof schema>;

export const RegisterForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: zodResolver(schema),
    defaultValues: {
      email: '',
      password: '',
      name: '',
      surname: '',
      patronymic: '',
      phone: '',
    },
  });

  const onSubmit = (data: FormData) => console.log(data);

  return (
    <form onSubmit={handleSubmit(onSubmit)} className='grid h-fit w-fit grid-cols-[1fr_1fr] grid-rows-[auto_auto] gap-5'>
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

      <div className='flex flex-col gap-5'>
        <Input
          {...register('name')}
          label='Имя'
          placeholder='Введите имя'
          className='h-12 w-[315px]'
          errorMessage={errors.name?.message}
          isInvalid={!!errors.name}
        />
        <Input
          {...register('surname')}
          label='Фамилия'
          placeholder='Введите фамилию'
          className='h-12 w-[315px]'
          errorMessage={errors.surname?.message}
          isInvalid={!!errors.surname}
        />
        <Input
          {...register('patronymic')}
          label='Отчество'
          placeholder='Введите отчество'
          className='h-12 w-[315px]'
          errorMessage={errors.patronymic?.message}
          isInvalid={!!errors.patronymic}
        />
      </div>
      <div className='col-span-2'>
        <Button type='submit' className='bg-uiDeepGray col-span-2 mt-4 p-2 text-white'>
          Зарегистрироваться
        </Button>
      </div>
    </form>
  );
};
