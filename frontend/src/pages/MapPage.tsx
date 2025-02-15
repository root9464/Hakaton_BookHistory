import { NavBar } from '@/components/NavBar';
import { useEffect, useRef, useState } from 'react';

type Point = { lat: number; lon: number };
export type PixelPosition = { x: number; y: number };

const project = (lat: number, lon: number, zoom: number) => {
  const scale = 1 << zoom;
  const worldSize = 256 * scale;
  const x = ((lon + 180) * worldSize) / 360;
  const siny = Math.sin((lat * Math.PI) / 180);
  const y = (0.5 - Math.log((1 + siny) / (1 - siny)) / (4 * Math.PI)) * worldSize;
  return { x, y };
};

const zoom = 10;
const center: Point = { lat: 52.3898, lon: 56.0525 };

const points: Point[] = [
  { lat: 52.335844, lon: 56.24909 },
  { lat: 52.278163, lon: 55.952936 },
  { lat: 52.509025, lon: 56.240041 },
];

export const useMarkerPositions = (containerRef: React.RefObject<HTMLElement | null>, points: Point[], center: Point, zoom: number) => {
  const [markers, setMarkers] = useState<PixelPosition[]>([]);

  useEffect(() => {
    const container = containerRef.current;
    if (!container) return;

    const calculateMarkers = () => {
      const rect = container.getBoundingClientRect();
      const containerCenterX = rect.width / 2;
      const containerCenterY = rect.height / 2;
      const centerPixel = project(center.lat, center.lon, zoom);

      const newMarkers = points.map((point) => {
        const pointPixel = project(point.lat, point.lon, zoom);
        return {
          x: containerCenterX + (pointPixel.x - centerPixel.x),
          y: containerCenterY + (pointPixel.y - centerPixel.y),
        };
      });

      setMarkers(newMarkers);
    };

    calculateMarkers();
    const observer = new ResizeObserver(calculateMarkers);
    observer.observe(container);

    return () => observer.disconnect();
  }, [containerRef, points, center, zoom]);

  return markers;
};

export default function MapPage() {
  const containerRef = useRef<HTMLDivElement>(null);

  const markers = useMarkerPositions(containerRef, points, center, zoom);

  const handleMarkerClick = (index: number) => {
    console.log(`Клик по маркеру ${index + 1}`, points[index]);
  };

  return (
    <div ref={containerRef} className='relative flex h-screen w-full items-center justify-center overflow-x-hidden'>
      <iframe
        src={`https://geois2.orb.ru/resource/8888/display/tiny?base=basemap_0&lon=56.0525&lat=52.3898&angle=0&zoom=${zoom}`}
        className='z-0 h-full w-full'
        style={{ pointerEvents: 'none' }}
      />

      {markers.map((position, index) => (
        <div
          key={index}
          className='absolute animate-pulse cursor-pointer'
          style={{
            left: position.x - 12,
            top: position.y - 12,
            width: '24px',
            height: '24px',
            background: 'radial-gradient(circle, #ff0000 40%, #ff000080 70%, #ff000020 100%)',
            borderRadius: '50%',
            zIndex: 10,
          }}
          onClick={() => handleMarkerClick(index)}
        />
      ))}

      <div className='absolute top-0 z-[1] h-max w-max bg-transparent'>
        <div className='relative h-full w-full bg-transparent px-2 pt-20 md:px-5'>
          <NavBar />
        </div>
      </div>
    </div>
  );
}
