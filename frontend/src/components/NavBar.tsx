import { UserLoginResponse } from '@/modules/auth/hooks/useAuth';
import { useQueryClient } from '@tanstack/react-query';
import { Link } from '@tanstack/react-router';

const PAGES = [
  {
    name: 'Главная',
    path: '/',
  },
  {
    name: 'История',
    path: '/history',
  },
  {
    name: 'Профиль',
    path: '/profile',
  },
];

export const NavBar = () => {
  const queryClient = useQueryClient();

  const cacheUserLoginData: UserLoginResponse | undefined = queryClient.getQueryData(['user']);

  return (
    <div className='m-auto flex h-fit w-fit flex-row items-center justify-between gap-5 rounded-[20px] bg-gradient-to-bl from-[#0F0D0D]/50 to-[#0F0D0D]/40 px-4 py-2.5 backdrop-blur-[15px]'>
      {PAGES.map(({ name, path }, index) => (
        <Link to={path} className='flex h-fit w-fit flex-row items-center justify-center bg-transparent text-white' key={index}>
          {name === 'Профиль' && cacheUserLoginData?.data ? (
            <div className='flex h-fit w-fit flex-row items-center justify-center bg-transparent text-white'>
              <img src='https://cdn-icons-png.flaticon.com/512/149/149071.png' className='mr-2 h-6 w-6 rounded-full' alt='' />
              {cacheUserLoginData?.data.name}
            </div>
          ) : (
            name
          )}
        </Link>
      ))}
    </div>
  );
};
