import { NavBar } from '@/components/NavBar';
import { AdminModule } from '@/modules/admin/Module';

export default function AdminPage() {
  return (
    <div className='relative flex h-screen w-full items-center justify-center overflow-hidden'>
      <div className='relative top-[calc(20px*4+44px)] h-full w-full px-2 py-5 md:top-[calc(20px*4+44px)] md:px-5'>
        <AdminModule />
      </div>

      <div className='absolute top-0 z-[1] h-max w-max bg-transparent'>
        <div className='relative h-full w-full bg-transparent px-2 pt-20 md:px-5'>
          <NavBar />
        </div>
      </div>
    </div>
  );
}
