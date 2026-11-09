'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { Mail, ShieldCheck, LogOut } from 'lucide-react';
import { useAuthStore } from '@/store/useAuthStore';

export default function ProfilePage() {
  const router = useRouter();
  const { user, logout } = useAuthStore();
  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);

  // Если не авторизован — сразу редиректим на /login
  useEffect(() => {
    if (mounted && !user) {
      router.replace('/login');
    }
  }, [mounted, user, router]);

  if (!mounted || !user) {
    return null;
  }

  const handleLogout = () => {
    logout();
    router.push('/');
  };

  return (
    <div className="py-12 px-4 max-w-md mx-auto">
      <div className="bg-white p-6 rounded-2xl border border-slate-200 shadow-sm text-center">
        {/* Аватар */}
        <div className="w-20 h-20 bg-emerald-100 text-emerald-700 font-bold text-2xl rounded-full flex items-center justify-center mx-auto mb-4 ring-4 ring-emerald-50">
          {user.name ? user.name.slice(0, 2).toUpperCase() : 'US'}
        </div>

        <h1 className="text-xl font-bold text-slate-900">
          {user.name || 'Пользователь'}
        </h1>
        <p className="text-sm text-slate-500 flex items-center justify-center gap-1.5 mt-1 mb-4">
          <Mail className="w-4 h-4 text-slate-400" />
          {user.email}
        </p>

        <span className="inline-flex items-center gap-1.5 px-3 py-1 rounded-full text-xs font-medium bg-slate-100 text-slate-700">
          <ShieldCheck className="w-3.5 h-3.5 text-emerald-600" />
          ID: #{user.id}
        </span>

        {/* Выход */}
        <div className="border-t border-slate-100 mt-6 pt-4">
          <button
            type="button"
            onClick={handleLogout}
            className="w-full flex items-center justify-center gap-2 text-sm font-semibold text-rose-600 hover:bg-rose-50 py-2.5 px-3 rounded-xl transition-colors cursor-pointer"
          >
            <LogOut className="w-4 h-4" />
            Выйти из аккаунта
          </button>
        </div>
      </div>
    </div>
  );
}