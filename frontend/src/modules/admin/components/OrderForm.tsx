import { Button, Input } from '@heroui/react';
import * as z from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { useOrder } from '../hooks/useOrder';

const OrderFormShema = z.object({
  name: z.string().min(1, 'Имя обязательно для заполнения'),
  files: z.instanceof(FileList).refine((files) => files.length > 0, 'Файл обязателен для загрузки'),
});

export type OrderFormData = z.infer<typeof OrderFormShema>;

export const OrderForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<OrderFormData>({
    resolver: zodResolver(OrderFormShema),
    defaultValues: {
      name: '',
      files: undefined, // FileList по умолчанию
    },
  });

  const { mutate } = useOrder();

  const onSubmit = (data: OrderFormData) => {
    const filesArray = Array.from(data.files); // Преобразуем FileList в массив
    const formData = {
      ...data,
      files: filesArray, // Заменяем FileList на массив файлов
    };
    mutate(data);
    
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className='flex flex-col gap-[30px]'>
      <Input
        {...register('name')}
        label='Name'
        placeholder='Введите Название'
        className='h-12 w-[315px]'
        errorMessage={errors.name?.message}
        isInvalid={!!errors.name}
      />
      <Input
        {...register('files')}
        label='Files'
        type='file'
        multiple // Разрешаем выбор нескольких файлов
        placeholder='Введите Файлы'
        className='h-12 w-[315px]'
        errorMessage={errors.files?.message}
        isInvalid={!!errors.files}
      />
      <Button type='submit' className='mt-4 bg-uiDeepGray p-2 text-white'>
        Добавить
      </Button>
    </form>
  );
};