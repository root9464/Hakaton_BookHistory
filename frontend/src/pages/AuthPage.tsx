import { AuthModule } from '../modules/auth/Module';

import SvoFrame from '@/assets//png/svo.png';

export default function AuthPage() {
  return (
    <div className='relative flex h-full w-full flex-row items-center justify-between pr-[40px]'>
      <AuthModule />

      <div className='relative flex h-fit w-fit flex-col items-center justify-center'>
        <p className='absolute bottom-5 text-3xl font-bold text-white'>Цени наших предков !</p>
        <img className='h-[600px] w-[500px] rounded-[40px]' src={SvoFrame} alt='svo frame' />
      </div>
    </div>
  );
}
