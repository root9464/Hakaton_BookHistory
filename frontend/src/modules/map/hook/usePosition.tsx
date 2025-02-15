import { RefObject, useEffect, useState } from 'react';
import { PixelPosition, Point } from '../Module';
import { project } from '../utils/utils';

export const useMarkerPositions = (containerRef: RefObject<HTMLElement | null>, points: Point[], center: Point, zoom: number) => {
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
