import { Admin } from './components/admin';

export const AdminModule = () => {
    return (
        <div className='relative flex h-screen w-[60%] flex-col items-center justify-center gap-5 rounded-r-[40px] bg-white'>
            <Admin />
        </div>
    );
};