'use client';
import React, { useContext, useEffect } from 'react';
import Heading from '../components/ui/Heading';
import { z } from 'zod';
import { InputValidated } from '../components/Form/InputValidated';
import FormValidated from '../components/Form/FormValidated';
import SubmitButton from '../components/Form/SubmitButton';
import { AuthenticationContext } from '../provider/AuthProvider';
import Link from 'next/link';
import { useRouter } from 'next/navigation';

const validator = z.object({
  email: z.string().email(),
  password: z.string().nonempty(),
});
const LoginPage = () => {
  const { login, user, error } = useContext(AuthenticationContext);
  const router = useRouter();

  useEffect(() => {
    if (user) {
      router.push('/');
    }
  }, [user, router]);

  const initialValues = {
    email: 'victorino4@mail.com',
    password: 'hola1234567',
  };
  return (
    <div className='app-content h-screen'>
      <div className='mx-auto flex h-full max-w-lg flex-col'>
        <div className='mx-auto flex w-full flex-col items-center justify-center px-6 py-8 md:h-screen lg:py-0'>
          <Heading className='my-5'>Login</Heading>
          {error && (
            <div className='rounded bg-red-500 p-2 text-white'>{error}</div>
          )}
          <div className='w-full flex-1'>
            <FormValidated
              validator={validator}
              initialValues={initialValues}
              onSubmit={(values) => {
                login(values);
              }}
            >
              <InputValidated
                name='email'
                type='email'
                label='Your email'
                className='w-full'
              />
              <InputValidated
                name='password'
                type='password'
                label='Password'
              />

              <SubmitButton
                label='login'
                className='block w-fit'
                size='lg'
                submittingMessage='loggin...'
              />
              <div>
                <p className='text-sm text-gray-500'>
                  Don&#39;t have an account yet?{' '}
                  <Link
                    href='/register'
                    className='font-bold text-primary-ligth'
                  >
                    Sign up
                  </Link>
                </p>
              </div>
            </FormValidated>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
