export const NavBar = () => (
  <div className='m-auto flex h-fit w-fit flex-row items-center justify-between gap-5 rounded-[20px] bg-gradient-to-bl from-[#0F0D0D]/50 to-[#0F0D0D]/40 px-4 py-2.5'>
    {['Главная', 'Профиль', 'История'].map((el, index) => (
      <div className='flex h-fit w-fit flex-row items-center justify-center bg-transparent text-white' key={index}>
        <p>{el}</p>
      </div>
    ))}
  </div>
);
