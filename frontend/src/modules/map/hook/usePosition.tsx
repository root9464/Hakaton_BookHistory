import { RefObject, useEffect, useState } from 'react';
import { PixelPosition, Point } from '../Module';

const project = (lat: number, lon: number, zoom: number) => {
  const scale = 1 << zoom;
  const worldSize = 256 * scale;
  const x = ((lon + 180) * worldSize) / 360;
  const siny = Math.sin((lat * Math.PI) / 180);
  const y = (0.5 - Math.log((1 + siny) / (1 - siny)) / (4 * Math.PI)) * worldSize;
  return { x, y };
};

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
