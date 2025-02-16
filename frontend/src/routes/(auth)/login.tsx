import OauthPage from '@/pages/OauthPage';
import { createFileRoute } from '@tanstack/react-router';

export const Route = createFileRoute('/(auth)/login')({
  component: OauthPage,
});
