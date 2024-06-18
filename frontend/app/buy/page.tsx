"use client";
import React, { useContext, useEffect, useState } from 'react'
import { AuthenticationContext } from '../provider/AuthProvider';
import Heading from '../components/ui/Heading';
import { buyBond, getBondsAvailable } from '../utils/api';
import { useMutation, useQuery } from '@tanstack/react-query';
import { Pagination } from '@mui/material';
import Table from '../components/ui/Table';
import Button from '../components/ui/Button';
import { formatMonetaryValue, formatNumericValue } from '../utils/index';
import { toast } from 'sonner';
import { ErrorWithRequest } from '../utils/Error';
import { BondWithOwner } from '../types';

const BuyBondPage = () => {
  const { user } = useContext(AuthenticationContext);
  const buyMutation = useMutation({ mutationFn: buyBond });
  const [page, setPage] = useState(1);
  const { data, status, refetch } = useQuery({
    queryKey: ['bonds', page],
    queryFn: getBondsAvailable,
    enabled: !!user,
  });
  const [bonds, setBonds] = useState<BondWithOwner[]>([]);

  useEffect(() => {
    if (data && data.bonds && data.bonds.length > 0) {
      setBonds(data.bonds);
    }
  }, [data]);

  function handleBuyClick(bondId:string) { 
    buyMutation.mutate(bondId, {
      onSuccess: (data) => {
        toast.success('Bond bought successfully');
        setBonds((prevBonds) => prevBonds.filter((bond) => bond.id !== bondId));
      },
      onError: (error) => {
        if (error instanceof ErrorWithRequest) {
          toast.error(error.message);
        }
        else {
          toast.error('An error occurred');
        }
      }
    });
  }
  return (
    <div className='app-content w-full h-full'>
      <Heading className='my-10 !text-primary'>
        Buy a bond
        </Heading>
        
      {
        status === 'pending' && <p>Loading...</p>
        }
      {
        status === 'success' && bonds.length > 0 && (
          <div>

            <Table
              head={
                <tr>
                  <Table.Head>ID</Table.Head>
                  <Table.Head>Name</Table.Head>
                  <Table.Head>Price</Table.Head>
                  <Table.Head>Currency</Table.Head>
                  <Table.Head>Number</Table.Head>
                  <Table.Head>Seller</Table.Head>
                  <Table.Head> </Table.Head>
                </tr>
              }
            >
              {
                bonds.map((bond) => (
                  <Table.Row key={bond.id}>
                    <Table.Data>{bond.id}</Table.Data>
                    <Table.Data>{bond.name}</Table.Data>
                    <Table.Data>{formatMonetaryValue(bond.price)}</Table.Data>
                    <Table.Data>{"MXN"}</Table.Data>
                    <Table.Data>{formatNumericValue(bond.number_bonds)}</Table.Data>
                    <Table.Data>{bond.owner}</Table.Data>
                    <Table.Data>
                      <Button
                        onClick={() => handleBuyClick(bond.id)}
                        
                      >
                        Buy
                      </Button>
                    </Table.Data>
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
          <div className='border rounded-md border-primary bg-primary-lighter px-4 py-4'>
            <p className='text-white'>

oh no! There are no bonds available to buy
            </p>
          </div>)

  }
    </div>
  )
}

export default BuyBondPage