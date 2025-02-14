import { createRouter, RouterProvider } from '@tanstack/react-router';

// Import the generated route tree
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { routeTree } from '../routeTree.gen';

// Create a new router instance
const router = createRouter({ routeTree });
const queryClient = new QueryClient();

// Register the router instance for type safety
declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router;
  }
}

export const GlobalProvider = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  );
};
