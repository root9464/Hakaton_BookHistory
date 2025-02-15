import { OrderModal } from './OrderModal';

export const SubBar = () => {
  return (
    <div className='m-auto flex h-fit w-fit flex-row items-center justify-between gap-5 rounded-[20px] bg-gradient-to-bl from-[#0F0D0D]/50 to-[#0F0D0D]/40 px-4 py-2.5 backdrop-blur-[15px]'>
      <OrderModal />
      <p className='flex h-fit w-fit flex-row items-center justify-center bg-transparent text-white'>fffffff</p>
    </div>
  );
};
