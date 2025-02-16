import { useWindow } from '@/shared/hooks/useWindow';
import { useDisclosure } from '@heroui/react';
import { useMemo, useRef, useState } from 'react';
import { Marker } from './components/Marker';
import { OrderModal } from './components/OrderModal';
import { useCords } from './hook/useCords';
import { useMarkerPositions } from './hook/usePosition';
import { project, toEPSG3857Direct } from './utils/utils';

export type Point = { lat: number; lon: number };
export type PixelPosition = { x: number; y: number };

const center: Point = { lat: 52.3898, lon: 56.0525 };

export const MapModule = () => {
  const containerRef = useRef<HTMLDivElement>(null);
  const { width, height } = useWindow();
  const zoom = width > height ? 10 : 9;
  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  const { data } = useCords();
  const [coordinates, setCoordinates] = useState('');

  const points = useMemo(() => {
    return data ? data.map((item) => ({ lat: item.cords.lat, lon: item.cords.lon, id: String(item.userID) })) : [];
  }, [data]);

  const markers = useMarkerPositions(containerRef, points, center, zoom);

  const handleClick = (event: React.MouseEvent<HTMLDivElement>) => {
    const container = containerRef.current;
    if (!container) return;

    const rect = container.getBoundingClientRect();
    if (event.clientX < rect.left || event.clientX > rect.right || event.clientY < rect.top || event.clientY > rect.bottom) return;

    const clickX = event.clientX - rect.left;
    const clickY = event.clientY - rect.top;

    const centerPixel = project(center.lat, center.lon, zoom);
    const offsetX = clickX - rect.width / 2;
    const offsetY = clickY - rect.height / 2;

    const worldX = centerPixel.x + offsetX;
    const worldY = centerPixel.y + offsetY;

    const { x, y } = toEPSG3857Direct(worldX, worldY, zoom);
    const stringCoord = `POINT (${x} ${y})`;
    console.log('Клик по карте', stringCoord);
    setCoordinates(stringCoord);
    onOpen();
  };

  return (
    <div className='relative h-full w-full'>
      <div className='h-full w-full' ref={containerRef} onClick={handleClick}>
        <iframe
          src={`https://geois2.orb.ru/resource/8888/display/tiny?base=basemap_0&lon=${center.lon}&lat=${center.lat}&angle=0&zoom=${zoom}`}
          className='pointer-events-none z-0 h-full w-full'
        />
        <Marker markers={markers} />
      </div>

      <OrderModal isOpen={isOpen} onOpen={onOpen} onOpenChange={onOpenChange} coordinates={coordinates} />
    </div>
  );
};
