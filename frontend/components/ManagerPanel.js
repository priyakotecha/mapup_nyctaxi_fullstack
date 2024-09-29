import { useEffect, useState } from 'react';
import { fetchManagerData } from '../services/api';
import RealTimeData from './RealTimeData';

const ManagerPanel = () => {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const res = await fetchManagerData();
      setData(res);
    };
    fetchData();
  }, []);

  return (
    <div className="manager-panel">
      <h2>Manager Panel</h2>
      <RealTimeData />
      <div className="data-list">
        {data.map((item) => (
          <div key={item.id}>{item.name}</div>
        ))}
      </div>
    </div>
  );
};

export default ManagerPanel;
