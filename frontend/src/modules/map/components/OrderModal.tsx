import { UserLoginResponse } from '@/modules/auth/hooks/useAuth';
import { Button, Input, Modal, ModalBody, ModalContent, ModalFooter, ModalHeader } from '@heroui/react';
import { zodResolver } from '@hookform/resolvers/zod';
import { useQueryClient } from '@tanstack/react-query';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { useOrder } from '../hook/useOrder';

const OrderFormSchema = z.object({
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

export type OrderFormData = z.infer<typeof OrderFormSchema>;

const fields = [
  { name: 'fio', label: 'ФИО', placeholder: 'Введите ФИО', type: 'text' },
  { name: 'date_of_birth', label: 'Дата рождения', placeholder: 'Введите дату рождения', type: 'date' },
  { name: 'place_of_birth', label: 'Место рождения', placeholder: 'Введите место рождения', type: 'date' },
  { name: 'name_of_millitary_commissariat', label: 'Название военкомата', placeholder: 'Введите название военкомата', type: 'text' },
  { name: 'millitary_rank', label: 'Воинское звание', placeholder: 'Введите воинское звание', type: 'text' },
  { name: 'date_of_death', label: 'Дата смерти', placeholder: 'Введите дату смерти', type: 'date' },
  { name: 'burial_place', label: 'Место захоронения', placeholder: 'Введите место захоронения', type: 'text' },
  { name: 'biographical_facts', label: 'Биографические факты', placeholder: 'Введите биографические факты', type: 'text' },
  { name: 'files', label: 'Файл', placeholder: 'Выберите файл', type: 'file' },
];

type OrderModalProps = {
  isOpen: boolean;
  onOpen: () => void;
  onOpenChange: () => void;
  coordinates: string;
};

export const OrderModal = ({ isOpen, onOpen, onOpenChange, coordinates }: OrderModalProps) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
    trigger,
    reset,
  } = useForm<OrderFormData>({
    resolver: zodResolver(OrderFormSchema),
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

  const queryClient = useQueryClient();

  const cacheUserLoginData: UserLoginResponse | undefined = queryClient.getQueryData(['user']);

  const { mutate } = useOrder();

  const onSubmit = (data: OrderFormData) => {
    const formData = new FormData();
    Object.entries({ ...data, sender_id: cacheUserLoginData?.data.id ?? '', geom: coordinates }).forEach(([key, value]) =>
      value instanceof FileList ? Array.from(value).forEach((file) => formData.append(key, file)) : formData.append(key, value as string),
    );
    console.log(data, 'data');
    reset();
    onOpenChange();
    mutate(formData);
  };

  return (
    <>
      <button onClick={onOpen} className='flex h-fit w-fit flex-row items-center justify-center bg-transparent text-white'>
        Create
      </button>
      <Modal isOpen={isOpen} onOpenChange={onOpenChange}>
        <ModalContent className='z-[2]'>
          {(onClose) => (
            <>
              <ModalHeader className='flex flex-col gap-1'>Создание записи</ModalHeader>
              <form onSubmit={handleSubmit(onSubmit)} className='flex flex-col gap-3'>
                <ModalBody>
                  {fields.map(({ name, label, placeholder, type }, index) => (
                    <Input
                      {...register(name as keyof OrderFormData, {
                        ...(type === 'file' && {
                          onChange: (e) => {
                            const file = e.target.files?.[0];
                            if (file) {
                              const dataTransfer = new DataTransfer();
                              dataTransfer.items.add(file);
                              e.target.files = dataTransfer.files;
                            }
                          },
                        }),
                      })}
                      key={index}
                      label={label}
                      placeholder={placeholder}
                      multiple={type === 'file'}
                      type={type}
                      className='h-fit w-full'
                      errorMessage={errors[name as keyof typeof errors]?.message}
                      isInvalid={!!errors[name as keyof typeof errors]}
                      onBlur={() => trigger(name as keyof OrderFormData)}
                    />
                  ))}
                </ModalBody>
                <ModalFooter>
                  <Button color='danger' variant='light' onPress={onClose} type='reset'>
                    Отмена
                  </Button>
                  <Button color='primary' type='submit'>
                    Создать
                  </Button>
                </ModalFooter>
              </form>
            </>
          )}
        </ModalContent>
      </Modal>
    </>
  );
};
