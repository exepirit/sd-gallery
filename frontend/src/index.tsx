import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { createTheme, NextUIProvider } from '@nextui-org/react';
import { Gallery } from './Features/Gallery/Gallery';
import { PictureViewContainer } from './Features/PictureView/Components';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

const router = createBrowserRouter([
  {
    path: "/",
    element: <Gallery />
  },
  {
    path: "/p/:pictureId",
    element: <PictureViewContainer />
  }
]);

const darkTheme = createTheme({
  type: 'light'
})

root.render(
  <React.StrictMode>
    <NextUIProvider theme={darkTheme}>
      <RouterProvider router={router}/>
    </NextUIProvider>
  </React.StrictMode>
);