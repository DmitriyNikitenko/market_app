'use client';

import { useEffect, useState } from 'react';
import Link from 'next/link';
import { ShoppingBag, User as UserIcon, LogOut } from 'lucide-react';
import { useCartStore } from '@/store/useCartStore';
import { useAuthStore } from '@/store/useAuthStore';

export default function Header() {
  const totalItems = useCartStore((state) => state.getTotalItems());
  const { user, logout } = useAuthStore();

  const [mounted, setMounted] = useState(false);
  useEffect(() => {
    setMounted(true);
  }, []);

  return (
    <header className="sticky top-0 z-50 bg-white/80 backdrop-blur-md border-b border-slate-200">
      <div className="max-w-6xl mx-auto px-4 sm:px-6 h-16 flex items-center justify-between">
        
        {/* Логотип */}
        <Link href="/" className="text-xl font-bold tracking-tight text-slate-900">
          Store<span className="text-emerald-600">.</span>
        </Link>

        {/* Навигация, Профиль и Корзина */}
        <div className="flex items-center space-x-5">
          <Link href="/products" className="text-sm font-medium text-slate-600 hover:text-slate-900 transition-colors">
            Каталог
          </Link>

          {/* Блок авторизации */}
          {mounted && user ? (
            <div className="flex items-center gap-3">
              {/* Ссылка на страницу профиля */}
              <Link
                href="/profile"
                className="text-xs font-semibold text-slate-700 bg-slate-100 hover:bg-slate-200 px-2.5 py-1 rounded-full transition-colors cursor-pointer"
                title="Перейти в профиль"
              >
                {user.name || user.email}
              </Link>

              <button
                type="button"
                onClick={logout}
                title="Выйти"
                className="text-slate-400 hover:text-rose-600 transition-colors p-1 cursor-pointer"
              >
                <LogOut className="w-4 h-4" />
              </button>
            </div>
          ) : (
            <Link
              href="/login"
              className="flex items-center gap-1.5 text-sm font-semibold text-slate-700 hover:text-emerald-600 transition-colors"
            >
              <UserIcon className="w-4 h-4" />
              Войти
            </Link>
          )}

          {/* Корзина */}
          <Link
            href="/cart"
            className="relative flex items-center justify-center p-2 rounded-full text-slate-700 hover:bg-slate-100 transition-colors"
            aria-label="Корзина"
          >
            <ShoppingBag className="w-6 h-6" />
            
            {mounted && totalItems > 0 && (
              <span className="absolute -top-1 -right-1 bg-emerald-600 text-white text-xs font-bold w-5 h-5 rounded-full flex items-center justify-center animate-in zoom-in-50 duration-200">
                {totalItems}
              </span>
            )}
          </Link>
        </div>

      </div>
    </header>
  );
}