import axios from 'axios';
import * as React from 'react';
import { useEffect, useState, useCallback } from 'react';

const Users = () => {
  const [users, setUsers] = useState<any>();
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');

  useEffect(() => {
    (async () => {
      const result = await axios.get('/users', { baseURL: process.env.BACKEND_URL });
      setUsers(result.data);
    })();
  }, []);

  const handleSubmit = useCallback(
    (ev: React.FormEvent<HTMLFormElement>) => {
      ev.preventDefault();
      (async () => {
        const result = await axios.post('/users', { firstName, lastName }, { baseURL: process.env.BACKEND_URL });
        setUsers((users: any) => [...users, result.data]);
      })();
    },
    [firstName, lastName]
  );

  return (
    <div>
      <h1>Users</h1>
      <table>
        {users?.map((user: any) => (
          <tr key={user._id}>
            <td>{user._id}</td>
            <td>{user.firstName}</td>
            <td>{user.lastName}</td>
            <td>{user.mail}</td>
          </tr>
        ))}
      </table>
      <form onSubmit={handleSubmit}>
        <label>First name</label>
        <input required value={firstName} onChange={ev => setFirstName(ev.target.value)} />
        <label>Last name</label>
        <input required value={lastName} onChange={ev => setLastName(ev.target.value)} />
        <button>Submit</button>
      </form>
    </div>
  );
};

export default Users;
