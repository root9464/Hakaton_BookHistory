import { HeroUIProvider } from '@heroui/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { createRouter, RouterProvider } from '@tanstack/react-router';
import { routeTree } from '../routeTree.gen';

const router = createRouter({
  routeTree,
  context: {
    userRole: null,
  },
});
const queryClient = new QueryClient();

declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router;
  }
}

export const GlobalProvider = () => {
  // хук для получения значения из store для контекста (лучше использовать store zustand а не чистый функционал tanstack router)
  return (
    <HeroUIProvider>
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router} />
      </QueryClientProvider>
    </HeroUIProvider>
  );
};
