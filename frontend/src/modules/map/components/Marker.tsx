import { useDisclosure } from '@heroui/react';
import { Fragment } from 'react/jsx-runtime';
import { PixelPosition } from '../Module';
import { AncestorModal } from './AncestorModal';

type MarkerProps = {
  markers: Array<PixelPosition & { id: string }>;
};

export const Marker = ({ markers }: MarkerProps) => {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  const handleMarkerClick = ({ x, y }: PixelPosition) => {
    console.log(`Клик по маркеру ${x} ${y}`);
  };

  return (
    <>
      {markers.map(({ id, x, y }, index) => (
        <Fragment key={index}>
          <div
            key={index}
            className='absolute z-[2] h-6 w-6 animate-pulse cursor-pointer rounded-full bg-red-500'
            style={{ left: x - 12, top: y - 12 }}
            onClick={(event) => {
              event.stopPropagation();
              onOpen();
              handleMarkerClick({ x, y });
            }}
          />
          <AncestorModal coordinates={{ id, lat: y, lon: x }} isOpen={isOpen} onOpenChange={onOpenChange} />
        </Fragment>
      ))}
    </>
  );
};
