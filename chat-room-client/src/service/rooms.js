import axios from 'axios';

export const getRooms = async (token) => {
  const response = await axios.get('http://localhost:9000/api/v1/get-room', {
    headers: {
      Authorization: `Bearer ${token}`, // Menambahkan header Authorization dengan token JWT
    },
  });
  console.log(response);
  return response.data;
};

export const createRoom = async (id, roomName, token) => {
  const response = await axios.post(
    'http://localhost:9000/api/v1/create-room',
    {
      id: id,
      name: roomName,
    },
    {
      headers: {
        Authorization: `Bearer ${token}`, // Menambahkan header Authorization dengan token JWT
      },
    }
  );
  console.log(response);
  return response.data;
};
