import React from 'react';
import Button from './Button';
import { AuthenticationContext } from '@/app/provider/AuthProvider';
import Heading from './Heading';
import Link from 'next/link';
import ButtonLink from './ButtonLink';
import { usePathname, useRouter } from 'next/navigation';

const Nabvar = () => {
  const { user, logout } = React.useContext(AuthenticationContext);
  const path = usePathname();
  if (path === '/login' || path === '/register') return null;
  if (!user)
    return (
      <div className='h-15 border border-primary px-4 py-2 font-bold text-white'>
        <nav className=''>
          <ul className='flex items-center justify-between gap-4'>
            <li>
              <Heading level={1} className=''>
                BondsApp
              </Heading>
            </li>
            <li>
              {' '}
              <ButtonLink color='dark' href='/register'>
                Create account
              </ButtonLink>
              <ButtonLink href='/login'>Login</ButtonLink>
            </li>
          </ul>
        </nav>
      </div>
    );
  return (
    <div className='h-15 bg-primary px-4 py-2 font-bold text-white'>
      <nav className=''>
        <ul className='flex items-center justify-between gap-4'>
          <li>
            <Link href='/'>
              <Heading level={1} className='!text-white'>
                BondsApp
              </Heading>
            </Link>
          </li>
          <li>
            <Link href={`/${user.username}/bonds`}>My bonds</Link>
          </li>
          <li>
            <Link href={`/buy`}>Bonds</Link>
          </li>
          <li>
            <Button onClick={() => logout()}>Logout</Button>
          </li>
        </ul>
      </nav>
    </div>
  );
};

export default Nabvar;
