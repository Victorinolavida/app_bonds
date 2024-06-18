'use client';

import classNames from 'classnames';

import React from 'react';

interface Props {
  level?: 1 | 2 | 3;
  className?: string;
  children: React.ReactNode;
}
const Heading = ({ level = 1, className, children }: Props) => {
  let defaultStyling = 'font-bold text-black ';
  switch (level) {
    case 1:
      defaultStyling += 'text-3xl';
      break;
    case 2:
      defaultStyling += 'text-2xl';
      break;
    default:
      defaultStyling += 'text-xl';
      break;
  }

  const HeadingTag = `h${level}` as keyof JSX.IntrinsicElements;
  return (
    <HeadingTag className={classNames(defaultStyling, className)}>
      {children}
    </HeadingTag>
  );
};

export default Heading;
