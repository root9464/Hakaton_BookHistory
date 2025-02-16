import { Button, Input } from '@heroui/react';
import { zodResolver } from '@hookform/resolvers/zod';
import { Link, useRouter } from '@tanstack/react-router';
import { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import * as z from 'zod';
import { useRegister } from '../hooks/useAuth';

const FormShema = z.object({
  email: z.string().email('Некорректный email'),
  password: z
    .string()
    .min(8, 'Пароль должен содержать минимум 8 символов')
    .regex(/[A-Z]/, 'Пароль должен содержать хотя бы одну заглавную букву')
    .regex(/[0-9]/, 'Пароль должен содержать хотя бы одну цифру'),
  name: z.string().min(1, 'Имя обязательно для заполнения'),
  surname: z.string().min(1, 'Фамилия обязательна для заполнения'),
  patronymic: z.string().optional(),
  phone: z
    .string()
    .min(10, 'Номер телефона должен содержать 10 цифр')
    .max(10, 'Номер телефона должен содержать 10 цифр')
    .regex(/^\d{10}$/, 'Номер телефона должен содержать только цифры')
    .transform((val) => `+7${val}`),
});

export type FormData = z.infer<typeof FormShema>;

const fields = [
  { name: 'email', label: 'Email', placeholder: 'Введите email', type: 'text' },
  { name: 'password', label: 'Password', placeholder: 'Придумайте пароль', type: 'password' },
  { name: 'phone', label: 'Phone', placeholder: 'Номер телефона', type: 'text' },

  { name: 'name', label: 'Имя', placeholder: 'Введите имя', type: 'text' },
  { name: 'surname', label: 'Фамилия', placeholder: 'Введите фамилию', type: 'text' },
  { name: 'patronymic', label: 'Отчество', placeholder: 'Введите отчество', type: 'text' },
];

export const RegisterForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: zodResolver(FormShema),
    defaultValues: {
      email: '',
      password: '',
      name: '',
      surname: '',
      patronymic: '',
      phone: '',
    },
  });
  const router = useRouter();

  const { data, mutate, isSuccess } = useRegister();

  useEffect(() => {
    if (isSuccess) {
      console.log(data);
      router.navigate({ to: '/auth' });
    }
  }, [isSuccess, data, router]);

  const onSubmit = (data: FormData) => mutate(data);

  return (
    <form onSubmit={handleSubmit(onSubmit)} className='grid h-fit w-fit grid-cols-[1fr_1fr] grid-rows-[auto_auto_auto] gap-5'>
      <div className='flex flex-col gap-5'>
        {fields.slice(0, 3).map(({ name, label, placeholder, type }, index) => (
          <Input
            {...register(name as keyof typeof register)}
            key={index}
            label={label}
            placeholder={placeholder}
            type={type}
            className='h-fit w-[315px]'
            errorMessage={errors[name as keyof typeof errors]?.message}
            isInvalid={!!errors[name as keyof typeof errors]}
          />
        ))}
      </div>

      <div className='flex flex-col gap-5'>
        {fields.slice(3).map(({ name, label, placeholder, type }, index) => (
          <Input
            {...register(name as keyof typeof register)}
            key={index}
            label={label}
            placeholder={placeholder}
            type={type}
            className='h-fit w-[315px]'
            errorMessage={errors[name as keyof typeof errors]?.message}
            isInvalid={!!errors[name as keyof typeof errors]}
          />
        ))}
      </div>

      <div className='col-span-2 flex w-1/2 flex-col items-center justify-center place-self-center'>
        <div className='flex flex-row gap-4 text-xs font-medium text-blue-600'>
          <Link to={'/auth'}>Уже есть аккаунт</Link>
          <Link to='.'>Войти через гос услуги</Link>
        </div>
        <Button type='submit' className='mt-4 bg-uiDeepGray p-2 text-white'>
          Зарегистрироваться
        </Button>
      </div>
    </form>
  );
};
