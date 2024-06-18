'use client';

import { ErrorMessage, Field } from 'formik';

interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label: string;
  name: string;
  className?: string;
}
export function InputValidated({ label, className, ...props }: InputProps) {
  return (
    <div>
      <label htmlFor={props.name}>
        {label}
        <Field {...props} className={className} />
        <ErrorMessage name={props.name} component='span' className='error' />
      </label>
    </div>
  );
}
