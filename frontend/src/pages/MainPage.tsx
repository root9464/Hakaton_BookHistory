import MainBg from '@/assets/png/bg.png';
import { MapMarker } from '@/assets/tsx/MapMarker';
import { NavBar } from '@/components/NavBar';
import { cn } from '@/shared/utils/utils';
import { Link } from '@tanstack/react-router';
import { HTMLMotionProps, motion } from 'framer-motion';
import { FC, useEffect, useState } from 'react';

const INTERESTING_FACTS = [
  `
    Потери после войны Вермахта составили около шести миллионов человек.
    По статистике, соотношение общего числа погибших и умерших людей между СССР и Германией составляет 7,3:1.
    Из этого делаем вывод что в СССР погибло более 43 млн людей. Эти цифры учитывают потери гражданских: СССР — 16,9 млн чел., Германия — 2 млн чел.
  `,

  `
    На протяжении десятилетия после Победы СССР формально находился еще в состоянии войны с Германией.
    После принятия капитуляции немцами СССР решил не принимать и не подписывать мир с врагом;
    и получается, что остался с Германией в состоянии войны.
  `,
];

export default function MainPage() {
  const [currentFactIndex, setCurrentFactIndex] = useState<number | null>(null);

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentFactIndex((prevIndex) => {
        let newIndex: number;
        do {
          newIndex = Math.floor(Math.random() * INTERESTING_FACTS.length);
        } while (newIndex === prevIndex);
        return newIndex;
      });
    }, 15 * 1000);

    return () => clearInterval(interval);
  }, []);

  return (
    <div className='relative flex h-screen w-full items-center justify-center overflow-x-hidden bg-red-700/35'>
      <img className='z-0 h-full w-full object-cover' src={MainBg} alt='svo frame' />

      <div className='absolute top-0 z-[1] h-full w-full bg-transparent'>
        <div className='relative h-full w-full bg-transparent px-5 pt-20'>
          <NavBar />
          <div className='relative mt-24 flex h-max w-fit flex-col items-center justify-center gap-5'>
            <Content
              className='w-[690px] px-4 py-2.5 text-white'
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.5 }}
            >
              <p>
                "Память сильнее времени" — это не просто проект, а дань уважения тем, кто отдал свои силы, жизнь и мужество во имя будущего. Мы
                сохраняем истории участников Великой Отечественной войны и современных героев, чтобы их подвиги не растворились в пыли времени, а
                остались в сердцах поколений.
              </p>
            </Content>
            <Content
              className='w-[690px] px-4 py-2.5 text-white'
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.5, delay: 0.2 }}
            >
              <p>
                Каждое имя, каждая судьба — это не просто строки на странице, а голос, звучащий сквозь десятилетия. Мы собираем воспоминания,
                письма, фотографии и свидетельства, чтобы передать их дальше — тем, кто ещё не родился, но должен знать, какой ценой завоёван
                мир. Наш проект — это тихая, но глубокая благодарность. Это память, которая живёт.
              </p>
            </Content>
            {currentFactIndex !== null && (
              <Content
                key={currentFactIndex}
                className='w-[690px] px-4 py-2.5 text-white'
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                exit={{ opacity: 0, y: 20 }}
                transition={{ duration: 0.5 }}
              >
                <p>{INTERESTING_FACTS[currentFactIndex]}</p>
              </Content>
            )}
          </div>

          {window.innerWidth >= 1980 && (
            <Link to='/register' className='absolute right-[25%] top-[56%] flex h-[100px] w-[100px] items-end justify-center bg-transparent'>
              <MapMarker className='h-10 w-10' />
            </Link>
          )}
        </div>
      </div>
    </div>
  );
}

const Content: FC<{ children: React.ReactNode; className?: string } & HTMLMotionProps<'div'>> = ({ children, className, ...props }) => {
  return (
    <motion.div className={cn(`rounded-[20px] bg-gradient-to-bl from-[#0F0D0D]/50 to-[#0F0D0D]/40 ${className}`)} {...props}>
      {children}
    </motion.div>
  );
};
