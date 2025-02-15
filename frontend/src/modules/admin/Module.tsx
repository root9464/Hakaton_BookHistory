import { AplicationForm } from "./components/aplicationForm";
import { OrderForm } from "./components/OrderForm";

export const AdminModule = () => {
  return (
    <div className='flex gap-[50px]'>
      <OrderForm />
      <AplicationForm />
    </div>
  );
};
