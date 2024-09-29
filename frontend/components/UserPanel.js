import { useEffect, useState } from 'react';
import { fetchUserData } from '../services/api';
import RealTimeData from './RealTimeData';

const UserPanel = () => {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const res = await fetchUserData();
      setData(res);
    };
    fetchData();
  }, []);

  return (
    <div className="user-panel">
      <h2>User Panel</h2>
      <RealTimeData />
      <div className="data-list">
        {data.map((item) => (
          <div key={item.id}>{item.name}</div>
        ))}
      </div>
    </div>
  );
};

export default UserPanel;
