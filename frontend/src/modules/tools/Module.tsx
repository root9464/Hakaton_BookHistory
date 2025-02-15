import { NavBar } from '@/components/NavBar';
import { SubBar } from './components/SubBar';

export const ToolsModule = () => {
  return (
    <div className='flex h-fit w-fit flex-row items-center justify-center gap-3'>
      <NavBar />
      <SubBar />
    </div>
  );
};
