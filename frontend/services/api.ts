import { useAuthStore } from '@/store/useAuthStore';

export const GATEWAY_URL =
  process.env.NEXT_PUBLIC_GATEWAY_URL || 'http://localhost:8080';

export async function gatewayFetch<T>(
  path: string,
  options: RequestInit = {}
): Promise<T> {
  const token = useAuthStore.getState().token;

  const response = await fetch(`${GATEWAY_URL}${path}`, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...options.headers,
    },
  });

  const data = await response.json().catch(() => ({}));

  if (!response.ok) {
    // Если токен протух или невалиден — сбрасываем сессию
    if (response.status === 401) {
      useAuthStore.getState().logout();
    }
    throw new Error(data.detail || data.message || 'Ошибка запроса к Gateway');
  }

  return data as T;
}