import { Header } from './components/Header';
import { RegisterForm } from './components/RegisterForm';

export const AuthModule = () => {
  return (
    <div className='relative flex h-screen w-[60%] flex-col items-center justify-center gap-5 rounded-r-[40px] bg-white'>
      <Header />
      <RegisterForm />
    </div>
  );
};
