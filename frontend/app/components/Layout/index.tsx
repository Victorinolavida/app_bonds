'use client';
import { AuthenticationContext } from '@/app/provider/AuthProvider';
import React, { useContext } from 'react';
import Nabvar from '../ui/Nabvar';

const Layout = ({ children }: { children: React.ReactNode }) => {
  const { user } = useContext(AuthenticationContext);

  return (
    <div>
      <Nabvar />
        <div className='flex flex-col gap-4 w-full h-full'>
          {children}
      </div>
    </div>
  );
};

export default Layout;
