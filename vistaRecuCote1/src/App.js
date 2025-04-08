import React, { useState, useEffect } from 'react';
import axios from 'axios';

const App = () => {
  const [users, setUsers] = useState([]);
  const [name, setName] = useState('');
  const [edad, setEdad] = useState(0);
  const [sexo, setSexo] = useState(true);
  const [count, setCount] = useState({ count_true: 0, count_false: 0 });

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const response = await axios.get('http://localhost:8080/user/longpoll');
        setUsers(response.data);
      } catch (error) {
        console.error('Error fetching users:', error);
      }
    };

    const fetchCount = () => {
      const eventSource = new EventSource('http://localhost:8080/user/poll/count');

      eventSource.onmessage = function(event) {
        try {
          const data = JSON.parse(event.data);
          if (data.count_true !== undefined && data.count_false !== undefined) {
            setCount(data);
          } else {
            console.log('Received non-count data: ', data);
          }
        } catch (error) {
          console.error('Error parsing event data:', error);
        }
      };

      eventSource.onerror = function(err) {
        console.error('EventSource error:', err);
        eventSource.close();
      };
    };

    fetchUsers();
    fetchCount();

    const intervalId = setInterval(() => {
      fetchUsers();
    }, 5000);

    return () => clearInterval(intervalId);
  }, []);

  const handleSubmit = async (event) => {
    event.preventDefault();
    const newUser = { edad, name, sexo };
    try {
      await axios.post('http://localhost:8080/user/', newUser);
      setName('');
      setEdad(0);
      setSexo(true);
    } catch (error) {
      console.error('Error creating user:', error);
    }
  };

  return (
    <div>
      <h1>Crear Usuario</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Nombre:</label>
          <input type="text" value={name} onChange={(e) => setName(e.target.value)} />
        </div>
        <div>
          <label>Edad:</label>
          <input type="number" value={edad} onChange={(e) => setEdad(Number(e.target.value))} />
        </div>
        <div>
          <label>Sexo:</label>
          <select value={sexo} onChange={(e) => setSexo(e.target.value === 'true')}>
            <option value="true">Hombre</option>
            <option value="false">Mujer</option>
          </select>
        </div>
        <button type="submit">Crear Usuario</button>
      </form>
      <h2>Usuarios</h2>
      <ul>
        {users && users.map((user) => (
          <li key={user.id}>{`${user.name} (${user.edad} años) - ${user.sexo ? 'Hombre' : 'Mujer'}`}</li>
        ))}
      </ul>
      <h2>Conteo de Usuarios</h2>
      <p>Hombres: {count.count_true}</p>
      <p>Mujeres: {count.count_false}</p>
    </div>
  );
};

export default App;