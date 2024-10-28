export interface Transaction {
    id: number;
    amount: number;
    date: string;
    type: 'entrada' | 'saida';
    userId: number;
    categoryId: number;
}
