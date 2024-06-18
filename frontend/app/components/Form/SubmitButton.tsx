import { useFormikContext } from 'formik';
import React from 'react';
import Button, { ButtonProps } from '../ui/Button';

interface Props extends ButtonProps {
  label: string;
  submittingMessage?: string;
}
const SubmitButton = ({ label, submittingMessage, ...props }: Props) => {
  const { isValid, isSubmitting, dirty, ...rest } = useFormikContext();
  return (
    <Button disabled={!(isValid || isSubmitting)} type='submit' {...props}>
      {isSubmitting ? submittingMessage : label}
    </Button>
  );
};

export default SubmitButton;
