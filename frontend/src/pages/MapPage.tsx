export default function MapPage() {
  return (
    <div className='relative flex h-screen w-full items-center justify-center'>
      <iframe
        src='https://geois2.orb.ru/resource/8888/display/tiny?base=basemap_0&lon=56.0525&lat=52.3898&angle=0&zoom=10&styles=8781%2C7986%2C7975%2C2092&linkMainMap=true&events=true&panel=none&controls=&panels='
        className='z-0 h-full w-full'
      />
    </div>
  );
}
