/* eslint-disable @typescript-eslint/no-unused-vars */
import { Spinner } from '@heroui/react';
import { useLayoutEffect } from 'react';

const ClientID = '33'; // Ваш ID в ЕЛК
const ClientSecret = 'J6fs2LlWwslyTq56jeuXfefXeoG3Mrh92PcSYe2z'; // Ваш секретный ключ
const RedirectURI = 'http://hackathon-5.orb.ru/profile/rsaag'; // Адрес редиректа
const AuthURL = 'https://lk.orb.ru/oauth/authorize'; // URL авторизации
const TokenURL = 'https://lk.orb.ru/oauth/token'; // URL для получения токенов
const UserInfoURL =
  'https://lk.orb.ru/api/get_user?scope=rsaag_id+personal_data+esia_data+email+phone+esia_user_id+organizations_user+auth_method'; // URL для получения информации о пользователе

const authURL = `${AuthURL}?client_id=${ClientID}&redirect_uri=${encodeURIComponent(RedirectURI)}&response_type=code&scope=email+auth_method&state=${encodeURIComponent('http://hackathon-5.orb.ru/')}`;

async function fetchTokens(code: string) {
  const response = await fetch(TokenURL, {
    method: 'POST',
    headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
    body: new URLSearchParams({
      client_id: ClientID,
      client_secret: ClientSecret,
      redirect_uri: RedirectURI,
      code,
      grant_type: 'authorization_code',
    }),
  });
  if (!response.ok) throw new Error('Ошибка получения токена');
  return response.json();
}

async function fetchUserInfo(accessToken: string) {
  const response = await fetch(UserInfoURL, {
    method: 'GET',
    headers: { Authorization: `Bearer ${accessToken}` },
  });
  if (!response.ok) throw new Error('Ошибка получения информации о пользователе');
  return response.json();
}

export default function OauthPage() {
  useLayoutEffect(() => {
    window.location.href = authURL;
  }, []);

  return (
    <div className='relative flex h-screen w-full items-center justify-center overflow-hidden'>
      <Spinner />
    </div>
  );
}
