import { User } from '@/store/useAuthStore';

const GATEWAY_URL = process.env.NEXT_PUBLIC_GATEWAY_URL || 'http://localhost:8080';

interface AuthResponse {
  access_token: string;
  token_type: string;
  user: User;
}

// 1. Вход в аккаунт
export async function loginUser(email: string, password: string): Promise<AuthResponse> {
  const response = await fetch(`${GATEWAY_URL}/api/auth/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password }),
  });

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}));
    throw new Error(errorData.detail || errorData.message || 'Ошибка входа в аккаунт');
  }

  return response.json();
}

// 2. Регистрация нового пользователя
export async function registerUser(email: string, password: string, name?: string): Promise<AuthResponse> {
  const response = await fetch(`${GATEWAY_URL}/api/auth/register`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password, name }),
  });

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}));
    throw new Error(errorData.detail || errorData.message || 'Ошибка регистрации');
  }

  return response.json();
}