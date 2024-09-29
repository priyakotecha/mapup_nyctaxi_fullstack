import React, { useState, useEffect } from 'react';
import { Bar } from 'react-chartjs-2';
import axios from 'axios';

function Dashboard() {
  const [chartData, setChartData] = useState({});

  useEffect(() => {
    const fetchData = async () => {
      const response = await axios.get('http://localhost:8080/getTaxiData');
      const data = response.data;

      const labels = data.map(d => d.pickup_datetime);
      const fares = data.map(d => d.fare_amount);

      setChartData({
        labels: labels,
        datasets: [
          {
            label: 'Fare Amount',
            data: fares,
            backgroundColor: 'rgba(75, 192, 192, 0.6)',
          }
        ]
      });
    };

    fetchData();
  }, []);

  return (
    <div>
      <h2>NYC Taxi Fare Distribution</h2>
      <Bar data={chartData} />
    </div>
  );
}

export default Dashboard;
``
