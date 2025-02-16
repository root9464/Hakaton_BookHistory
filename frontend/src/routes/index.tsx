import { createFileRoute, redirect } from '@tanstack/react-router';
import MainPage from '../pages/MainPage';

export const Route = createFileRoute('/')({
  component: MainPage,
  beforeLoad: async ({ context }) => {
    if (context.userRole !== 'user') {
      throw redirect({ to: '/auth' });
    }
  },
});
