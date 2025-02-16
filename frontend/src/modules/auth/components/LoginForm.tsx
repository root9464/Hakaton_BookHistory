import { Button, Input } from '@heroui/react';
import { zodResolver } from '@hookform/resolvers/zod';
import { Link, useRouter } from '@tanstack/react-router';
import { useAtom } from 'jotai';
import { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import * as z from 'zod';
import { useLogin } from '../hooks/useAuth';
import { UserRoleAtom } from '../store/userRole';

const FormShema = z.object({
  email: z.string().email('Некорректный email'),
  password: z
    .string()
    .min(8, 'Пароль должен содержать минимум 8 символов')
    .regex(/[A-Z]/, 'Пароль должен содержать хотя бы одну заглавную букву')
    .regex(/[0-9]/, 'Пароль должен содержать хотя бы одну цифру'),
});

export type LoginFormData = z.infer<typeof FormShema>;

const fields = [
  { name: 'email', label: 'Email', placeholder: 'Введите email', type: 'text' },
  { name: 'password', label: 'Password', placeholder: 'Придумайте пароль', type: 'password' },
];

export const LoginForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormData>({
    resolver: zodResolver(FormShema),
    defaultValues: {
      email: '',
      password: '',
    },
  });
  const router = useRouter();
  const [, setRole] = useAtom(UserRoleAtom);
  const { data, mutate, isSuccess } = useLogin();

  useEffect(() => {
    if (isSuccess) {
      console.log(data);
      router.navigate({ to: '/' });
      setRole(data.data.role);
    }
  }, [isSuccess, data, router, setRole]);

  const onSubmit = (data: LoginFormData) => mutate(data);
  return (
    <form onSubmit={handleSubmit(onSubmit)} className='grid h-fit w-fit grid-cols-[1fr_1fr] grid-rows-[auto_auto_auto] gap-5'>
      {fields.map(({ name, label, placeholder, type }, index) => (
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

      <div className='col-span-2 flex w-1/2 flex-col items-center justify-center place-self-center'>
        <div className='flex flex-row gap-4 text-xs font-medium text-blue-600'>
          <Link to={'/register'}>Пока нету аккаунта</Link>
          <Link to='/'>Войти через гос услуги</Link>
        </div>
        <Button type='submit' className='mt-4 bg-uiDeepGray p-2 text-white'>
          Зарегистрироваться
        </Button>
      </div>
    </form>
  );
};
