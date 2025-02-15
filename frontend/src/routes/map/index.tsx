import MapPage from '@/pages/MapPage';
import { createFileRoute } from '@tanstack/react-router';

export const Route = createFileRoute('/map/')({
  component: MapPage,
});
