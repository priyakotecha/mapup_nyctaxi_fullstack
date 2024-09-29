import { useEffect, useState } from 'react';
import { fetchStreamData } from '../services/api';  // Use API call for fetching stream data

const RealTimeData = () => {
  const [streamData, setStreamData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const res = await fetchStreamData();
      setStreamData(res);
    };

    // Polling the API every 5 seconds to get real-time data
    const intervalId = setInterval(() => {
      fetchData();
    }, 5000);

    return () => clearInterval(intervalId); // Cleanup
  }, []);

  return (
    <div className="real-time-data">
      <h3>Real-Time Data</h3>
      {streamData.map((data, index) => (
        <div key={index}>{data}</div>
      ))}
    </div>
  );
};

export default RealTimeData;
