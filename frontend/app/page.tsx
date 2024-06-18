'use client';

import { useContext,  useState } from 'react';
import { AuthenticationContext } from './provider/AuthProvider';
import LandingPage from './components/LandingPage';
import { useQuery } from '@tanstack/react-query';
import { getUserOwnedBonds } from './utils/api';
import Table from './components/ui/Table';
import { Pagination } from '@mui/material';
import Heading from './components/ui/Heading';

export default function Home({ children }: { children: React.ReactNode }) {
  const { user } = useContext(AuthenticationContext);
  const [page, setPage] = useState(1);
  const {data, status, refetch} = useQuery({
    queryKey: ['bonds', page],
    queryFn: getUserOwnedBonds,
    enabled: !!user,
    staleTime: 1000 * 60 ,
    
  });
  if (!user) return <LandingPage />;
  return (
    <main className='app-content w-full h-full'>
      <Heading className='my-10 !text-primary'>Welcome to the BondsApp <strong>
        {user.username}
      </strong>
      </Heading>
      <div>
        
        {
          status === 'pending' && <p>Loading...</p>
      }
        {
           status === 'success' && data && data.bonds  && data.bonds.length >0 && (
            <div>

            <Table
              head={
                <tr>
                  <Table.Head>ID</Table.Head>
                  <Table.Head>Name</Table.Head>
                  <Table.Head>Currency</Table.Head>
                  <Table.Head>Number</Table.Head>
                  <Table.Head>Status</Table.Head>
                </tr>
              }
            >
              {
                data.bonds.map((bond) => (
                  <Table.Row key={bond.id}>
                    <Table.Data>{bond.id}</Table.Data>
                    <Table.Data>{bond.price}</Table.Data>
                    <Table.Data>{"MXN"}</Table.Data>
                    <Table.Data>{bond.number_bonds}</Table.Data>
                    <Table.Data>{bond.status}</Table.Data>
                    </Table.Row>
                ))
              }

              </Table>
              <Pagination
                color='primary'
              className='mt-4 w-full flex justify-start'
                count={data.pagination.last_page}
                page={data.pagination.current_page}
                onChange={(event, page) => {
                  setPage(page);
                  refetch();
                }}
              />
            </div>
              )
      }
        {
          status === 'success' && data && data.bonds && data.bonds.length === 0 && (
            <p>No bonds found</p>)

        }
      </div>
</main>
  );
}
