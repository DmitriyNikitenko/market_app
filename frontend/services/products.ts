import { gatewayFetch } from './api';

export interface Product {
  id: number;
  title: string;
  category: string;
  price: number;
  image: string;
}

export async function getProducts(): Promise<Product[]> {
  const raw = await gatewayFetch<any>('/api/menu', { method: 'GET' });
  const list = Array.isArray(raw) ? raw : raw.content ?? [];

  return list.map((item: any) => ({
    id: item.id,
    title: item.title ?? item.name,
    category: item.category,
    price: Number(item.price),
    image: item.image,
  }));
}