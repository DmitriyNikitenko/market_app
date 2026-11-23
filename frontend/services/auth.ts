import { User } from '@/store/useAuthStore';
import { gatewayFetch } from './api';

export interface AuthResponse {
  token: string;
  user: User;
}

interface TokenResponse {
  token: string;
}

// Получение профиля пользователя по токену из user_service
export function fetchCurrentUser(token: string): Promise<User> {
  return gatewayFetch<User>('/api/users/me', {
    method: 'GET',
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
}


export function firstFetchCurrentUser(
  firstName: string,
  token: string
): Promise<User> {
  return gatewayFetch<User>('/api/users/me', {
    method: 'PUT',
    headers: {
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify({ firstName }),
  });
}


// 1. Вход в аккаунт
export async function loginUser(
  email: string,
  password: string
): Promise<AuthResponse> {
  const { token } = await gatewayFetch<TokenResponse>('/api/auth/login', {
    method: 'POST',
    body: JSON.stringify({ email, password }),
  });

  const user = await fetchCurrentUser(token);
  return { token, user };
}

// 2. Регистрация нового пользователя
export async function registerUser(
  email: string,
  password: string,
  firstName: string
): Promise<AuthResponse> {
  const { token } = await gatewayFetch<TokenResponse>('/api/auth/register', {
    method: 'POST',
    body: JSON.stringify({ email, password }),
  });
  const user = await firstFetchCurrentUser(firstName, token);
  return { token, user };
}