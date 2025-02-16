import { useQuery } from '@tanstack/react-query';

const MOCK = {
  data: {
    id: '0e73f9f0-345f-425e-b3ae-b3357df10bae',
    sender_id: 'f79a75eb-2c99-459d-b5d1-c14730f7cff0',
    fio: 'Иванов Иван Иванович',
    geom: 'POINT (6259532.948015831 6873896.547813708)', // Координаты примерные
    date_of_birth: '1915-05-10',
    place_of_birth: 'Москва, СССР',
    name_of_millitary_commissariat: 'Московский военный комиссариат',
    millitary_rank: 'Капитан',
    date_of_death: '1943-02-15',
    burial_place: 'Братская могила, Волгоградская область',
    biographical_facts:
      'Участвовал в Великой Отечественной войне. Отличился в битве под Сталинградом, командовал взводом. Награждён Орденом Красного Знамени.',
    status: 'confirmed',
    files: [
      {
        id: '7ef9921d-3239-49c8-bebb-662cf36b0434',
        name: 'photo_ivanov_ww2.png',
        owner_id: '0e73f9f0-345f-425e-b3ae-b3357df10bae',
        owner_type: 'applications',
      },
    ],
    rewards: [
      {
        id: '8b1db63f-3472-48ca-9622-8b6285787401',
        name: 'Орден Красного Знамени',
        image: {
          id: 'reward_img_1',
          name: 'order_of_red_banner.png',
          owner_id: '0e73f9f0-345f-425e-b3ae-b3357df10bae',
          owner_type: 'rewards',
        },
      },
      {
        id: '9c2df63a-4571-48db-89a3-1c63b784d002',
        name: 'Медаль За отвагу',
        image: {
          id: 'reward_img_2',
          name: 'medal_for_courage.png',
          owner_id: '0e73f9f0-345f-425e-b3ae-b3357df10bae',
          owner_type: 'rewards',
        },
      },
    ],
  },
  message: 'Application found successfully',
  status: 'success',
};

export const useAncestor = (id: string) =>
  useQuery({
    queryKey: ['ancestors', id],
    queryFn: async () => {
      const { data } = (await new Promise((resolve) => {
        setTimeout(() => {
          resolve(MOCK);
        }, 500);
      })) as typeof MOCK;

      return data;
    },

    staleTime: 1000 * 60,
  });
