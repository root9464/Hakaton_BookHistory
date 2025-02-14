import { Outlet, createRootRoute } from '@tanstack/react-router';
import * as React from 'react';

export const Route = createRootRoute({
  component: Layout,
});

function Layout() {
  return (
    <React.Fragment>
      <Outlet />
    </React.Fragment>
  );
}
