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
  username: z.string().min(4).max(50),
  password: z.string().min(8).max(50),
});
const RegisterPage = () => {
  const { register, user, error } = useContext(AuthenticationContext);
  const router = useRouter();
  const initialValues = {
    username: '',
    email: '',
    password: '',
  };
  return (
    <div className='app-content h-screen'>
      <div className='mx-auto flex h-full max-w-lg flex-col'>
        <div className='mx-auto flex w-full flex-col items-center justify-center px-6 py-8 md:h-screen lg:py-0'>
          <Heading className='my-5'>Sign In</Heading>
          {error && (
            <div className='rounded bg-red-500 p-2 text-white'>{error}</div>
          )}
          <div className='w-full flex-1'>
            <FormValidated
              validator={validator}
              initialValues={initialValues}
              onSubmit={(values) => {
                register(values);
              }}
            >
              <InputValidated
                name='email'
                type='email'
                label='Your email'
                className='w-full'
              />
              <InputValidated
                name='username'
                type='text'
                label='Your Username'
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
                submittingMessage='Creating account...'
              />
              <div>
                <p className='text-sm text-gray-500'>
                  Do you have an account?{' '}
                  <Link href='/login' className='font-bold text-primary-light'>
                    login
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

export default RegisterPage;
