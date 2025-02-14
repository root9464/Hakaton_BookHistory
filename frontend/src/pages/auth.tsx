import soldier from "@/assets/soldier.png";
import { useForm } from "react-hook-form";

export const Auth = () => {
  const { register } = useForm();

  return (
    <div className="bg-[#DADBE0] w-full h-screen flex justify-between">
      <div className="bg-white flex flex-col text-center justify-center items-center w-[960px] rounded-tr-[40px] rounded-br-[40px] drop-shadow-sm">
        <h1><b>Вход</b></h1>
        <p>– Почти память о своих предках и вспомни историю</p>
        <form className="flex flex-col w-[318px] h-max">
          <input
            type="text"
            placeholder="Email"
            className="w-full h-[46px] rounded-[12px] bg-[#F4F4F5] border-none outline-none text-[14px] pl-[12px] mt-[37px]"
            {...register("email")}
          />
          <input
            type="text"
            placeholder="Password"
            className="w-full h-[46px] rounded-[12px] bg-[#F4F4F5] border-none outline-none text-[14px] pl-[12px] mt-[37px]"
            {...register("password")}
          />
          <input
            type="text"
            placeholder="Number"
            className="w-full h-[46px] rounded-[12px] bg-[#F4F4F5] border-none outline-none text-[14px] pl-[12px] mt-[37px]"
            {...register("number")}
          />
          <a href="" className="text-[12px] text-[#001AFF] mt-[37px]">Войти через гос услуги</a>
          <button className="w-[318px] h-[40px] rounded-[12px] bg-[#373B40] text-white mt-[7px]">Зарегистрироваться</button>
        </form>
      </div>
      <div className="flex justify-center items-center w-[960px] h-full">
        <img src={soldier} alt="" className="max-w-[736px] max-h-[897px] min-h-[400px] min-w-[350px]" />
      </div>
    </div>
  );
};
