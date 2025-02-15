import { atom } from 'jotai';

type UserRole = 'user' | 'admin' | null;

const userRoleState = atom<UserRole>(null);

const UserRoleAtom = atom(
  (get) => get(userRoleState),
  (_, set, role: UserRole) => set(userRoleState, role),
);

export { UserRoleAtom, type UserRole };
