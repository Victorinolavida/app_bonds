'use client';

import { useContext, useState } from 'react';
import { AuthenticationContext } from './provider/AuthProvider';
import LandingPage from './components/LandingPage';
import { useMutation, useQuery } from '@tanstack/react-query';
import { buyBond, getUserOwnedBonds } from './utils/api';
import Table from './components/ui/Table';
import { Pagination } from '@mui/material';
import Heading from './components/ui/Heading';
import { formatMonetaryValue, formatNumericValue } from './utils/index';
import ButtonLink from './components/ui/ButtonLink';

export default function Home({ children }: { children: React.ReactNode }) {
  const { user } = useContext(AuthenticationContext);
  const [page, setPage] = useState(1);
  const [openModal, setOpenModal] = useState(false);
  const { data, status, refetch } = useQuery({
    queryKey: ['bonds', page],
    queryFn: getUserOwnedBonds,
    enabled: !!user,
  });

  if (!user) return <LandingPage />;
  return (
    <main className='app-content h-full w-full'>

      <Heading className='my-10 !text-primary'>
        Welcome to the BondsApp <strong>{user.username}</strong>
      </Heading>
      <div className='flex md:justify-between justify-end'>
      <ButtonLink 
      href='/new'
      color="dark"
      >Create a new bond</ButtonLink>
        </div>
      <div>
        {status === 'pending' && <p>Loading...</p>}
        {status === 'success' &&
          data &&
          data.bonds &&
          data.bonds.length > 0 && (
            <div>
              <Table
                head={
                  <tr>
                    <Table.Head>ID</Table.Head>
                    <Table.Head>Name</Table.Head>
                    <Table.Head>Price</Table.Head>
                    <Table.Head>Currency</Table.Head>
                    <Table.Head>Number</Table.Head>
                    <Table.Head>Status</Table.Head>
                  </tr>
                }
              >
                {data.bonds.map((bond) => (
                  <Table.Row key={bond.id}>
                    <Table.Data>{bond.id}</Table.Data>
                    <Table.Data>{bond.name}</Table.Data>
                    <Table.Data>{formatMonetaryValue(bond.price)}</Table.Data>
                    <Table.Data>{'MXN'}</Table.Data>
                    <Table.Data>
                      {formatNumericValue(bond.number_bonds)}
                    </Table.Data>
                    <Table.Data>{bond.status}</Table.Data>
                  </Table.Row>
                ))}
              </Table>
              <Pagination
                color='primary'
                className='mt-4 flex w-full justify-start'
                count={data.pagination.last_page}
                page={data.pagination.current_page}
                onChange={(event, page) => {
                  setPage(page);
                  refetch();
                }}
              />
            </div>
          )}
        {status === 'success' &&
          data &&
          data.bonds &&
          data.bonds.length === 0 && (
            <div className='rounded-md border border-primary bg-primary-lighter px-4 py-4'>
              <p className='text-white'>You do not have any bonds yet.</p>
            </div>
          )}
      </div>
    </main>
  );
}
