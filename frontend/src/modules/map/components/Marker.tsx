import { PixelPosition } from '../Module';

type MarkerProps = {
  markers: PixelPosition[];
  handleMarkerClick: (index: number) => void;
};

export const Marker = ({ handleMarkerClick, markers }: MarkerProps) => {
  return (
    <>
      {markers.map((position, index) => (
        <div
          key={index}
          className='absolute z-[1] h-6 w-6 animate-pulse cursor-pointer rounded-full bg-red-500'
          style={{ left: position.x - 12, top: position.y - 12 }}
          onClick={() => handleMarkerClick(index)}
        />
      ))}
    </>
  );
};
