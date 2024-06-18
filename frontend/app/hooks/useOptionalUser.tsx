'use client';
import { useRouter } from 'next/navigation';
import React, { useContext } from 'react';
import { AuthenticationContext } from '../provider/AuthProvider';
import { User } from '../types';

export const useOptionalUser = (): User => {
  const router = useRouter();

  const { user } = useContext(AuthenticationContext);
  console.log(user);
  if (!user) {
    // @ts-ignore
    return router.push('/login');
  }
  return user;
};
