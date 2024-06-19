import React from 'react';
import Heading from '../ui/Heading';
import Image from 'next/image';

const LandingPage = () => {
  return (
    <main className='app-content h-full w-full'>
      <Heading className='my-10 !text-primary'>
        Welcome to the BondsApp
      </Heading>
      <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Minus facere illo consequatur saepe aspernatur hic nulla sequi expedita recusandae! Placeat assumenda expedita explicabo dolore odio ut delectus, eveniet nisi aut.</p>
      <div>
        <Image src='/stocks.jpg' alt='bonds' width={500} height={300} />
      </div>

    </main>
  );
};

export default LandingPage;
