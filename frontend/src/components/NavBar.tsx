import { Link } from '@tanstack/react-router';

const PAGES = [
  {
    name: 'Главная',
    path: '/',
  },
  {
    name: 'Профиль',
    path: '/profile',
  },
  {
    name: 'История',
    path: '/history',
  },
];

export const NavBar = () => (
  <div className='m-auto flex h-fit w-fit flex-row items-center justify-between gap-5 rounded-[20px] bg-gradient-to-bl from-[#0F0D0D]/50 to-[#0F0D0D]/40 px-4 py-2.5'>
    {PAGES.map(({ name, path }, index) => (
      <Link to={path} className='flex h-fit w-fit flex-row items-center justify-center bg-transparent text-white' key={index}>
        {name}
      </Link>
    ))}
  </div>
);
