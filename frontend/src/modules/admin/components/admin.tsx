// import { useAtom } from 'jotai';
// import { usersAtom, applicationsAtom, User, Application } from '@/store/atom';
// import { v4 as uuidv4 } from 'uuid';

export const Admin = () => {
  // const [users, setUsers] = useAtom(usersAtom);
  // const [applications, setApplications] = useAtom(applicationsAtom);

  // //добавление пользователя
  // const addUser = () => {
  //   const newUser: User = {
  //     id: uuidv4(),
  //     name: 'Иван',
  //     surname: 'Иванов',
  //     patronymic: 'Иванович',
  //     email: 'ivan@example.com',
  //     password: 'password123',
  //     phone: '1234567890',
  //   };
  //   setUsers([...users, newUser]);
  // };

  //добавление заявки
  // const addApplication = (userId: string) => {
  //   const newApplication: Application = {
  //     id: uuidv4(),
  //     userId,
  //     status: 'pending',
  //     description: 'Новая заявка',
  //   };
  //   setApplications([...applications, newApplication]);
  // };

  // //удаление пользователя и его заявок
  // const deleteUser = (userId: string) => {
  //   // Удаляем пользователя
  //   const updatedUsers = users.filter((user) => user.id !== userId);
  //   setUsers(updatedUsers);

  //   // удаление все заявки
  //   const delitedApplications = applications.filter((app) => app.userId !== userId);
  //   setApplications(delitedApplications);
  // };

  // //смены статуса
  // const toggleApplicationStatus = (applicationId: string) => {
  //   const updatedApplications = applications.map((app) =>
  //     app.id === applicationId ? { ...app, status: app.status === 'pending' ? ('approved' as const) : ('pending' as const) } : app,
  //   );
  //   setApplications(updatedApplications);
  // };

  // //удаление заявки (отдельно от пользователя (хз зачем))
  // const deleteApplication = (applicationId: string) => {
  //   const updatedApplications = applications.filter((app) => app.id !== applicationId);
  //   setApplications(updatedApplications);
  // };


};
