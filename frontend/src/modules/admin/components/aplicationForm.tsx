import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { Button, Input } from '@heroui/react';

// Схема валидации
const FormSchema = z.object({
  fio: z.string().min(1, 'ФИО обязательно для заполнения'),
  date_of_birth: z.string().min(1, 'Дата рождения обязательна для заполнения'),
  place_of_birth: z.string().min(1, 'Место рождения обязательно для заполнения'),
  name_of_millitary_commissariat: z.string().min(1, 'Название военкомата обязательно для заполнения'),
  millitary_rank: z.string().min(1, 'Воинское звание обязательно для заполнения'),
  date_of_death: z.string().min(1, 'Дата смерти обязательна для заполнения'),
  burial_place: z.string().min(1, 'Место захоронения обязательно для заполнения'),
  biographical_facts: z.string().min(1, 'Биографические факты обязательны для заполнения'),
  files: z.instanceof(FileList).refine((files) => files.length > 0, 'Файл обязателен для загрузки'),
});

// Тип данных формы
export type FormData = z.infer<typeof FormSchema>;

export const AplicationForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: zodResolver(FormSchema), // Подключаем валидацию через zod
    defaultValues: {
      fio: '',
      date_of_birth: '',
      place_of_birth: '',
      name_of_millitary_commissariat: '',
      millitary_rank: '',
      date_of_death: '',
      burial_place: '',
      biographical_facts: '',
      files: undefined,
    },
  });


  const onSubmit = (data: FormData) => {
    const filesArray = data.files ? Array.from(data.files) : [];

    const multipleFormData = {
      ...data, 
      files: filesArray,
    };

    console.log('Форма отправлена:', multipleFormData);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className='flex flex-col gap-[15px]'>
      {/* Поле fio */}
      <Input
        {...register('fio')}
        label='Fio'
        placeholder='Введите fio'
        className='h-12 w-[315px]'
        errorMessage={errors.fio?.message}
        isInvalid={!!errors.fio}
      />

      {/* Поле Дата рождения */}
      <Input
        {...register('date_of_birth')}
        label='Дата рождения'
        placeholder='Введите дату рождения'
        className='h-12 w-[315px]'
        errorMessage={errors.date_of_birth?.message}
        isInvalid={!!errors.date_of_birth}
      />

      {/* Поле Место рождения */}
      <Input
        {...register('place_of_birth')}
        label='Место рождения'
        placeholder='Введите место рождения'
        className='h-12 w-[315px]'
        errorMessage={errors.place_of_birth?.message}
        isInvalid={!!errors.place_of_birth}
      />

      {/* Поле Название военкомата */}
      <Input
        {...register('name_of_millitary_commissariat')}
        label='Название военкомата'
        placeholder='Введите название военкомата'
        className='h-12 w-[315px]'
        errorMessage={errors.name_of_millitary_commissariat?.message}
        isInvalid={!!errors.name_of_millitary_commissariat}
      />

      {/* Поле Воинское звание */}
      <Input
        {...register('millitary_rank')}
        label='Воинское звание'
        placeholder='Введите воинское звание'
        className='h-12 w-[315px]'
        errorMessage={errors.millitary_rank?.message}
        isInvalid={!!errors.millitary_rank}
      />

      {/* Поле Дата смерти */}
      <Input
        {...register('date_of_death')}
        label='Дата смерти'
        placeholder='Введите дату смерти'
        className='h-12 w-[315px]'
        errorMessage={errors.date_of_death?.message}
        isInvalid={!!errors.date_of_death}
      />

      {/* Поле Место захоронения */}
      <Input
        {...register('burial_place')}
        label='Место захоронения'
        placeholder='Введите место захоронения'
        className='h-12 w-[315px]'
        errorMessage={errors.burial_place?.message}
        isInvalid={!!errors.burial_place}
      />

      {/* Поле Биографические факты */}
      <Input
        {...register('biographical_facts')}
        label='Биографические факты'
        placeholder='Введите биографические факты'
        className='h-12 w-[315px]'
        errorMessage={errors.biographical_facts?.message}
        isInvalid={!!errors.biographical_facts}
      />

      {/* Поле Файлы */}
      <Input
        {...register('files')}
        label='Файлы'
        multiple
        type='file'
        className='h-12 w-[315px]'
        errorMessage={errors.files?.message}
        isInvalid={!!errors.files}
      />

      {/* Кнопка отправки формы */}
      <Button type='submit' className='mt-4 bg-uiDeepGray p-2 text-white'>
        Отправить
      </Button>
    </form>
  );
};
