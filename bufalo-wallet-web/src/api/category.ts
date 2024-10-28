import api from './api';
import { Category } from '../types/category';

export const getCategories = async (): Promise<Category[]> => {
    const response = await api.get<Category[]>('/categories');
    return response.data;
};

export const createCategory = async (category: Category): Promise<Category> => {
    const response = await api.post<Category>('/categories', category);
    return response.data;
};
