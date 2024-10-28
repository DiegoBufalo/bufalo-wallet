import api from './api';
import { User } from '../types/user';

export const getUsers = async (): Promise<User[]> => {
    const response = await api.get<User[]>('/users');
    return response.data;
};

export const createUser = async (user: User): Promise<User> => {
    const response = await api.post<User>('/users', user);
    return response.data;
};
