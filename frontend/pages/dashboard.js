import { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import AdminPanel from '../components/AdminPanel';
import ManagerPanel from '../components/ManagerPanel';
import UserPanel from '../components/UserPanel';
import { getRole, logout } from '../utils/auth';

const Dashboard = () => {
  const [role, setRole] = useState('');
  const router = useRouter();

  useEffect(() => {
    // Get role from local storage
    const userRole = getRole();
    if (!userRole) {
      router.push('/login');
    } else {
      setRole(userRole);
    }
  }, []);

  const handleLogout = () => {
    logout();
    router.push('/login');
  };

  return (
    <div className="dashboard-container">
      <h1>Dashboard</h1>
      <button onClick={handleLogout}>Logout</button>

      {role === 'admin' && <AdminPanel />}
      {role === 'manager' && <ManagerPanel />}
      {role === 'user' && <UserPanel />}
    </div>
  );
};

export default Dashboard;
