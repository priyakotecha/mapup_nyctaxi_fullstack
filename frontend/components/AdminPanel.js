import { useEffect, useState } from 'react';
import { fetchAllData } from '../services/api';
import RealTimeData from './RealTimeData';

const AdminPanel = () => {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const res = await fetchAllData();
      setData(res);
    };
    fetchData();
  }, []);

  return (
    <div className="admin-panel">
      <h2>Admin Panel</h2>
      <RealTimeData />
      <div className="data-list">
        {data.map((item) => (
          <div key={item.id}>{item.name}</div>
        ))}
      </div>
    </div>
  );
};

export default AdminPanel;
