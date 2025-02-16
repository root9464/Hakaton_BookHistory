import { Spinner } from '@heroui/react';
import { useQuery } from '@tanstack/react-query';
import { useRouter } from '@tanstack/react-router';
import axios from 'axios';
import { useEffect, useLayoutEffect } from 'react';

export default function OauthPage() {
  const router = useRouter();
  useLayoutEffect(() => {
    window.location.href =
      'https://lk.orb.ru/oauth/authorize?client_id=33&redirect_uri=http://hackathon-5.orb.ru/profile/rsaag&response_type=code&scope=email+auth_method&state=http://hackathon-5.orb.ru/';
  }, []);

  const { data, isSuccess } = useQuery({
    queryKey: ['oauth'],
    queryFn: async () => {
      const { data, status, statusText } = await axios.get(`http://hackathon-5.orb.ru/user/data`);
      if (status !== 200) {
        throw new Error(`${status}: ${statusText}`);
      }

      return data;
    },
  });

  useEffect(() => {
    if (isSuccess && data) {
      router.navigate({ to: '/map' });
    }
  }, [data, isSuccess, router]);

  return (
    <div className='relative flex h-screen w-full items-center justify-center overflow-hidden'>
      <Spinner />
    </div>
  );
}
