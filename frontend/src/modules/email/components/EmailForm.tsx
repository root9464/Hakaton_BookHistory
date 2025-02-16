import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as z from 'zod';
import { Button } from '@heroui/react';
import { Textarea } from '@heroui/input';
import { useEmail } from '../hooks/useEmail';

const EmailShema = z.object({
  description: z.string().min(1, 'Имя обязательно для заполнения'),
});

export type EmailData = z.infer<typeof EmailShema>;

export const EmailForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<EmailData>({
    resolver: zodResolver(EmailShema),
    defaultValues: {
      description: '',
    },
  });

  const { data, mutate } = useEmail();
  console.log(data);

  const onSubmit = (data: EmailData) => mutate(data);

  return (
    <form onSubmit={handleSubmit(onSubmit)} className='flex flex-col gap-5'>
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
