import { useEffect, useState } from 'react';
import { fromEvent } from 'rxjs';
import { debounceTime, distinctUntilChanged, map } from 'rxjs/operators';

export const useWindow = () => {
  const [size, setSize] = useState({
    width: window.innerWidth,
    height: window.innerHeight,
  });

  useEffect(() => {
    const resize$ = fromEvent(window, 'resize').pipe(
      debounceTime(100),
      map(() => ({
        width: window.innerWidth,
        height: window.innerHeight,
      })),
      distinctUntilChanged((prev, curr) => prev.width === curr.width && prev.height === curr.height),
    );

    const subscription = resize$.subscribe(setSize);

    return () => subscription.unsubscribe();
  }, []);

  return size;
};
