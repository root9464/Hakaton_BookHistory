import { Spinner } from '@heroui/react';
import { useLayoutEffect } from 'react';

export default function OauthPage() {
  useLayoutEffect(() => {
    window.location.href =
      'https://lk.orb.ru/oauth/authorize?client_id=33&redirect_uri=http://hackathon-5.orb.ru/profile/rsaag&response_type=code&scope=email+auth_method&state=http://hackathon-5.orb.ru/';
  }, []);

  return (
    <div className='relative flex h-screen w-full items-center justify-center overflow-hidden'>
      <Spinner />
    </div>
  );
}
