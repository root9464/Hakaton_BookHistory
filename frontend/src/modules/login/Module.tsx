import { Header } from './components/Header';
import { LoginForm } from './components/LoginForm';

export const LoginModule = () => {
  return (
    <div className='relative flex h-screen w-[60%] flex-col items-center justify-center gap-5 rounded-r-[40px] bg-white'>
      <Header />
      <LoginForm />
    </div>
  );
};
