import classNames from 'classnames';
import Link from 'next/link';
import React from 'react';

type buttonColors = 'red' | 'dark' | 'purple' | 'blue';
type buttonSize = 'xs' | 'sm' | 'base' | 'lg' | 'xl';
export interface ButtonProps
  extends React.AnchorHTMLAttributes<HTMLAnchorElement> {
  href: string;
  size?: buttonSize;
  color?: buttonColors;
}
const ButtonLink = ({
  children,
  color = 'blue',
  size = 'base',
  href,
  className,
  ...props
}: ButtonProps) => {
  let styling = '';
  switch (color) {
    case 'red':
      styling +=
        'focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm me-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900';
      break;
    case 'dark':
      styling +=
        'text-white bg-gray-800 hover:bg-gray-900 focus:outline-none focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm me-2 mb-2 dark:bg-gray-800 dark:hover:bg-gray-700 dark:focus:ring-gray-700 dark:border-gray-700';
      break;
    case 'purple':
      styling +=
        'focus:outline-none text-white bg-purple-700 hover:bg-purple-800 focus:ring-4 focus:ring-purple-300 font-medium rounded-lg text-sm mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900';
      break;
    default:
      styling +=
        'text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800';
      break;
  }
  switch (size) {
    case 'xs':
      styling += ' px-3 py-2 text-xs font-medium';
      break;
    case 'sm':
      styling += ' px-3 py-2 text-sm font-medium';
      break;

    case 'lg':
      styling += ' px-5 py-3 text-base font-medium';
      break;
    case 'xl':
      styling += ' px-6 py-3.5 text-base font-medium ';
      break;
    default:
      styling += ' px-4 py-2.5 text-base font-medium ';
      break;
  }
  return (
    <Link href={href} className={classNames(className, styling)} {...props}>
      {children}
    </Link>
  );
};

export default ButtonLink;
