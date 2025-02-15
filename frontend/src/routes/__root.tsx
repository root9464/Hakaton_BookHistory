import { UserRole } from '@/modules/auth/store/userRole';
import { Outlet, createRootRouteWithContext } from '@tanstack/react-router';
import * as React from 'react';

export const Route = createRootRouteWithContext<{ userRole: UserRole | null }>()({
  component: Layout,
});

function Layout() {
  return (
    <React.Fragment>
      <Outlet />
    </React.Fragment>
  );
}
