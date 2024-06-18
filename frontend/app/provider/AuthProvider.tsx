'use client';

import { useMutation, useQuery } from '@tanstack/react-query';
import React, { useEffect } from 'react';
import {
  userLogin,
  userLogout,
  userRegister,
  validateSession,
} from '../utils/api';
import { User } from '../types';
import { usePathname, useRouter } from 'next/navigation';
import { ErrorWithRequest } from '../utils/Error';
import { Toaster } from 'sonner';

export const AuthenticationContext = React.createContext<{
  login: (credentials: Record<string, unknown>) => void;
  register: (credentials: Record<string, unknown>) => void;
  logout: () => void;
  user: User | null;
  error: string | null;
}>({
  login: () => {},
  logout: () => {},
  register: () => {},
  user: null,
  error: null,
});

const publicPages = ['/login', '/register', '/'];
export default function AuthProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const { data, status, isFetching } = useQuery({
    queryKey: ['session'],
    queryFn: validateSession,
  });
  const loginMutation = useMutation({ mutationFn: userLogin });
  const logoutMutation = useMutation({ mutationFn: userLogout });
  const registerMutation = useMutation({ mutationFn: userRegister });
  const [user, setUser] = React.useState<User | null>(null);
  const [error, setError] = React.useState<string | null>(null);
  const router = useRouter();
  const path = usePathname();

  useEffect(() => {
    if (status === 'success' && !isFetching && data?.user) {
      setUser({
        id: data.user.id,
        username: data.user.username,
      });
    }
  }, [data, status, isFetching]);

  useEffect(() => {

    if (user && publicPages.includes(path)) {
      router.push('/');
    }
    

  }, [user, path, router]);

  const login = function (data: Record<string, unknown>) {
    loginMutation.mutate(data, {
      onSuccess: (data) => {
        setUser({
          username: data.user.username,
          id: data.user.id,
        });
        setError(null);
        router.push('/');
      },
      onError: (error) => {
        if (error instanceof ErrorWithRequest) {
          setError(error.message);
          console.log(error);
        } else {
          setError('An error occurred');
        }
      },
    });
  };

  const register = function (data: Record<string, unknown>) {
    registerMutation.mutate(data, {
      onSuccess: (data) => {
        setUser({
          username: data.user.username,
          id: data.user.id,
        });
        setError(null);
        router.push('/');
      },
      onError: (error) => {
        if (error instanceof ErrorWithRequest) {
          setError(error.message);
        } else {
          setError('An error occurred');
        }
      },
    });
  };

  const logout = async () => {
    logoutMutation.mutate(null, {
      onSuccess: (data) => {
        setUser(null);
        router.push('/login');
      },
      onError: (error) => {
        console.log(error);
      },
    });
  };

  return (
    <AuthenticationContext.Provider
      value={{ user, login, logout, error, register }}
    >
      <Toaster richColors position="top-right"/>
      {children}
    </AuthenticationContext.Provider>
  );
}
