import { MapModule } from '@/modules/map/Module';
import { ToolsModule } from '@/modules/tools/Module';

export default function MapPage() {
  return (
    <div className='relative flex h-screen w-full items-center justify-center overflow-hidden'>
      <MapModule />

      <div className='absolute top-0 z-[1] h-max w-max bg-transparent'>
        <div className='relative h-full w-full bg-transparent px-2 pt-20 md:px-5'>
          <ToolsModule />
        </div>
      </div>
    </div>
  );
}
