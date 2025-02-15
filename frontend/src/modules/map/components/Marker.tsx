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
    </>
  );
};
