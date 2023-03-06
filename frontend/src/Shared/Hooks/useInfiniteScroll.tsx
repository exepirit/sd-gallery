import { useState, useEffect, Dispatch, SetStateAction } from 'react';

export const useInfiniteScroll = (callback: () => void): [boolean, Dispatch<SetStateAction<boolean>>] => {
  const [isFetching, setIsFetching] = useState(false);

  useEffect(() => {
    if (window.innerHeight + document.documentElement.scrollTop >= document.documentElement.offsetHeight)
      setIsFetching(true);
    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  useEffect(() => {
    if (!isFetching) return;
    callback();
  }, [isFetching]);

  function handleScroll() {
    if (window.innerHeight + document.documentElement.scrollTop < document.documentElement.offsetHeight || isFetching) return;
    setIsFetching(true);
  }

  return [isFetching, setIsFetching];
}