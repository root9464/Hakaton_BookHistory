import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as z from 'zod';
import { Input, Button } from '@heroui/react';
import { Textarea } from '@heroui/input';

const FormShema = z.object({
  email: z.string().email('Некорректный email'),
  description: z.string().min(1, 'Имя обязательно для заполнения'),
});

type FormData = z.infer<typeof FormShema>;

export const EmailForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: zodResolver(FormShema),
    defaultValues: {
      email: '',
      description: '',
    },
  });

  const onSubmit = (data: FormData) => console.log(data);

  return (
    <form onSubmit={handleSubmit(onSubmit)} className='flex flex-col gap-5'>
      <Input
        {...register('email')}
        label='Email'
        placeholder='Введите email'
        className='h-12 w-[315px]'
        color='warning'
        errorMessage={errors.email?.message}
        isInvalid={!!errors.email}
      />
      <Textarea
        className='max-w-xs'
        label='Сообщение'
        placeholder='Введите сообщение'
        {...register('description')}
        errorMessage={errors.description?.message}
        isInvalid={!!errors.description}
        color='warning'
      />
      <Button className='bg-gradient-to-tr from-pink-500 to-yellow-500 text-white shadow-lg' radius='full' type='submit'>
        Отправить
      </Button>
    </form>
  );
};
