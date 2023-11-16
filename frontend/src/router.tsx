import { Suspense, lazy } from 'react';
import { Navigate } from 'react-router-dom';
import { RouteObject } from 'react-router';

import SidebarLayout from 'src/layouts/SidebarLayout';
import SuspenseLoader from 'src/components/SuspenseLoader';

const Loader = (Component) => (props) =>
  (
    <Suspense fallback={<SuspenseLoader />}>
      <Component {...props} />
    </Suspense>
  );

// Dashboards
const Settings = Loader(lazy(() => import('src/content/dashboards/Settings')));

// Applications
const Status404 = Loader(
  lazy(() => import('src/content/pages/Status/Status404'))
);

const routes: RouteObject[] = [
  {
    path: '',
    element: <SidebarLayout />,
    children: [
      {
        path: '/',
        element: <Navigate to="dashboards" replace />
      },
      {
        path: 'dashboards',
        element: <Settings />
      },
      {
        path: '*',
        element: <Status404 />
      }
    ]
  }
];

export default routes;
