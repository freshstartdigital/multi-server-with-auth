import React, { FC, useEffect, useState } from 'react';

const HomePageLayout: FC<{ text: string }> = ({ text }) => {
  const [state, setState] = useState<string>('');

  const fetchText = async () => {
    try {
      const url = process.env.NODE_ENV === 'production' ? '/api/hello' : 'http://localhost:8080/hello';
      const res = await fetch(url, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          
        },

      });

      if (!res.ok) {
        throw new Error(`Error: ${res.status}`);
      }

      const data = await res.json();
      console.log(data);
      setState(data.text);
    } catch (error) {
      console.error('Failed to fetch data:', error);
    }
  };

  useEffect(() => {
    fetchText();
  }, []);

  return (
    <div>
      {text}
      This is fetched text: {state}
    </div>
  );
};

export default HomePageLayout;
