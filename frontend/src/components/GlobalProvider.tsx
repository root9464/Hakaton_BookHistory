import { UserRoleAtom } from '@/modules/auth/store/userRole';
import { HeroUIProvider } from '@heroui/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import { createRouter, RouterProvider } from '@tanstack/react-router';
import { useAtom } from 'jotai';
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
  const [userRole] = useAtom(UserRoleAtom);
  return (
    <HeroUIProvider>
      <QueryClientProvider client={queryClient}>
        <RouterProvider
          router={router}
          context={{
            userRole,
          }}
        />

        <ReactQueryDevtools initialIsOpen={false} />
      </QueryClientProvider>
    </HeroUIProvider>
  );
};
