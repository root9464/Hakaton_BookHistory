import { NavBar } from '@/components/NavBar';
import { MapModule } from '@/modules/map/Module';

export default function MapPage() {
  return (
    <div className='relative flex h-screen w-full items-center justify-center overflow-x-hidden'>
      <MapModule />

      <div className='absolute top-0 z-[1] h-max w-max bg-transparent'>
        <div className='relative h-full w-full bg-transparent px-2 pt-20 md:px-5'>
          <NavBar />
        </div>
      </div>
    </div>
  );
}
