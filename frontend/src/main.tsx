import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { GlobalProvider } from './components/GlobalProvider.tsx';
import './index.css';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <GlobalProvider />
  </StrictMode>,
);
