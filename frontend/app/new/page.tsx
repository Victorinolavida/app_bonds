"use client"
import React from 'react'


import Heading from '../components/ui/Heading';
import FormValidated from '../components/Form/FormValidated';
import { z } from 'zod';
import { useMutation } from '@tanstack/react-query';
import { createBond } from '../utils/api';
import { MAX_BOND_PRICE_AS_INTEGER } from '../utils/constants';
import { InputValidated } from '../components/Form/InputValidated';
import SubmitButton from '../components/Form/SubmitButton';
import { toast } from 'sonner';
import { ErrorWithRequest } from '../utils/Error';
const validatorNewBond = z.object({
    price: z.number().gte(0, 'Price must be greater than 0').lte(MAX_BOND_PRICE_AS_INTEGER, 'Price must be less than 100,000,000'),
    number: z.number().min(1, 'Number must be greater than 1').max(10_000, 'Number must be less than 10,000'),
    name: z.string().min(3, 'Name must be at least 3 characters long').max(40, 'Name must be less than 40 characters long'),
});
const CreateBoundPage = () => {
    const createBoundMutation = useMutation({mutationFn: createBond });
  return (
<div className='app-content h-screen'>
          <div className='mx-auto flex h-full max-w-lg flex-col gap-4'>
<Heading level={1}>Create a new Bond</Heading>
            <FormValidated
            initialValues={{ name: '', price: '', number: '' }}
                validator={validatorNewBond}
                onSubmit={(values) => {
                    try {
                        
                        const data = validatorNewBond.parse(values);

                        createBoundMutation.mutate(data, {
                            onSuccess: () => {
                                toast.success('Bond created successfully');
                            },
                            onError: (error) => {
                                if (error instanceof ErrorWithRequest) {
                                   return  toast.error(error.message);
                                }
                                toast.error('An error occurred');

                            }
                        });


                    } catch (error) {
                        if (error instanceof z.ZodError) {
                            console.log(error)
                        }
                    }
                }}
                >
                <InputValidated name="name" label="Name" />
                <InputValidated name="price" label="Price" type="number"/>
                <InputValidated name="number" label="Number" type="number"/>
                <SubmitButton label="Create" submittingMessage='submitting..'/>
            </FormValidated>
      </div>

      </div>
  )
}

export default CreateBoundPage