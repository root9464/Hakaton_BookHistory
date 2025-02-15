import { NavBar } from '@/components/NavBar';

export default function MapPage() {
  return (
    <div className='relative flex h-screen w-full items-center justify-center overflow-x-hidden'>
      <iframe
        src='https://geois2.orb.ru/resource/8888/display/tiny?base=basemap_0&lon=56.0525&lat=52.3898&angle=0&zoom=10&styles=8781%2C7986%2C7975%2C2092&linkMainMap=true&events=false&panel=none&controls=&panels='
        className='z-0 h-full w-full'
      />

      <div className='absolute top-0 z-[1] h-max w-full bg-transparent'>
        <div className='relative h-full w-full bg-transparent px-2 pt-20 md:px-5'>
          <NavBar />
        </div>
      </div>
    </div>
  );
}
