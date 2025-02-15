import { useRef } from 'react';
import { Marker } from './components/Marker';
import { useMarkerPositions } from './hook/usePosition';

export type Point = { lat: number; lon: number };
export type PixelPosition = { x: number; y: number };

const zoom = 10;
const center: Point = { lat: 52.3898, lon: 56.0525 };

const points: Point[] = [
  { lat: 52.335844, lon: 56.24909 },
  { lat: 52.278163, lon: 55.952936 },
  { lat: 52.277743, lon: 56.196009 },
];

export const MapModule = () => {
  const containerRef = useRef<HTMLDivElement>(null);

  const markers = useMarkerPositions(containerRef, points, center, zoom);

  const handleMarkerClick = (index: number) => {
    console.log(`Клик по маркеру ${index + 1}`, points[index]);
  };
  return (
    <div className='relative h-full w-full' ref={containerRef}>
      <iframe
        src={`https://geois2.orb.ru/resource/8888/display/tiny?base=basemap_0&lon=56.0525&lat=52.3898&angle=0&zoom=${zoom}`}
        className='z-0 h-full w-full'
        style={{ pointerEvents: 'none' }}
      />
      <Marker markers={markers} handleMarkerClick={handleMarkerClick} />
    </div>
  );
};
