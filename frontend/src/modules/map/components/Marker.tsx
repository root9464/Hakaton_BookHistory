import { useDisclosure } from '@heroui/react';
import { PixelPosition } from '../Module';
import { AncestorModal } from './AncestorModal';

type MarkerProps = {
  markers: Array<PixelPosition & { id: string }>;
};

export const Marker = ({ markers }: MarkerProps) => {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();

  return (
    <>
      {markers.map(({ id, x, y }, index) => (
        <AncestorModal coordinates={{ id, lat: y, lon: x }} isOpen={isOpen} onOpenChange={onOpenChange} key={index}>
          <button
            type='button'
            className='absolute z-[1] h-6 w-6 cursor-pointer rounded-full bg-red-500'
            style={{ left: x - 12, top: y - 12 }}
            onClick={(event) => {
              event.stopPropagation();
              onOpen();
            }}
          />
        </AncestorModal>
      ))}
    </>
  );
};
