import { EmailForm } from "../email/components/EmailForm";
import { AplicationForm } from "./components/aplicationForm";
import { OrderForm } from "./components/OrderForm";

export const AdminModule = () => {
  return (
    <div className='flex gap-[50px] pt-10'>
      <OrderForm />
      <AplicationForm />
      <EmailForm />
    </div>
  );
};
