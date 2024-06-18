'use client';
import classNames from 'classnames';
import { ErrorMessage, Field, Form, Formik } from 'formik';
import React from 'react';
import { set, z } from 'zod';
import { toFormikValidate } from 'zod-formik-adapter';

interface Props {
  validator: z.AnyZodObject;
  initialValues: z.infer<Props['validator']>;
  children: React.ReactNode;
  onSubmit: (values: Props['initialValues']) => void;
  className?: string;
}

const FormValidated = ({
  validator,
  initialValues,
  children,
  onSubmit,
}: Props) => {
  return (
    <Formik
      initialValues={initialValues}
      validate={toFormikValidate(validator)}
      onSubmit={(values, { setSubmitting }) => {
        onSubmit(values);
        setSubmitting(false);
      }}
    >
      {() => (
        <Form className={classNames('flex flex-col gap-4', classNames)}>
          {children}
        </Form>
      )}
    </Formik>
  );
};

export default FormValidated;
