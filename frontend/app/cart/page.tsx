'use client';

import { useEffect, useState } from 'react';
import Link from 'next/link';
import { Trash2, Plus, Minus, ArrowLeft, ShoppingBag } from 'lucide-react';
import { useCartStore } from '@/store/useCartStore';
import PromoCodeInput from '@/components/PromoCodeInput';

export default function CartPage() {
  const {
    items,
    updateQuantity,
    removeFromCart,
    clearCart,
    getTotalPrice,
    getRawTotalPrice,
    discountPercent,
    promoCode,
  } = useCartStore();

  // Защита от расхождения серверного и клиентского HTML при чтении localStorage
  const [mounted, setMounted] = useState(false);
  useEffect(() => {
    setMounted(true);
  }, []);

  if (!mounted) {
    return (
      <div className="max-w-4xl mx-auto py-16 px-4 text-center">
        <div className="h-8 w-48 bg-slate-200 rounded animate-pulse mx-auto mb-4" />
        <div className="h-32 bg-slate-100 rounded-xl animate-pulse" />
      </div>
    );
  }

  // Состояние: пустая корзина
  if (items.length === 0) {
    return (
      <div className="max-w-3xl mx-auto py-20 px-4 text-center">
        <div className="w-20 h-20 bg-slate-100 text-slate-400 rounded-full flex items-center justify-center mx-auto mb-6">
          <ShoppingBag className="w-10 h-10" />
        </div>
        <h1 className="text-2xl font-bold text-slate-900 mb-2">Ваша корзина пуста</h1>
        <p className="text-slate-500 mb-8">
          Перейдите в каталог, чтобы выбрать понравившиеся товары.
        </p>
        <Link
          href="/"
          className="inline-flex items-center gap-2 bg-emerald-600 text-white px-6 py-3 rounded-lg font-semibold hover:bg-emerald-500 transition-colors"
        >
          <ArrowLeft className="w-4 h-4" />
          Вернуться к покупкам
        </Link>
      </div>
    );
  }

  const rawTotalPrice = getRawTotalPrice();
  const totalPrice = getTotalPrice();
  const discountAmount = rawTotalPrice - totalPrice;

  return (
    <div className="py-10 px-4 sm:px-6 lg:px-8 max-w-6xl mx-auto">
      <div className="flex items-center justify-between mb-8 border-b border-slate-200 pb-5">
        <div>
          <h1 className="text-3xl font-extrabold text-slate-900">Корзина</h1>
          <p className="text-sm text-slate-500 mt-1">Всего позиций: {items.length}</p>
        </div>
        <button
          onClick={clearCart}
          className="text-sm text-rose-600 hover:text-rose-700 font-medium transition-colors"
        >
          Очистить корзину
        </button>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        {/* Список товаров (2 колонки) */}
        <div className="lg:col-span-2 space-y-4">
          {items.map((item) => (
            <div
              key={item.id}
              className="flex items-center gap-4 bg-white p-4 rounded-xl border border-slate-200 shadow-sm"
            >
              {/* Фото */}
              <div className="w-20 h-20 bg-slate-100 rounded-lg overflow-hidden shrink-0">
                <img
                  src={item.image}
                  alt={item.title}
                  className="w-full h-full object-cover"
                />
              </div>

              {/* Название и цена за 1 шт. */}
              <div className="flex-1 min-w-0">
                <h2 className="text-sm font-semibold text-slate-900 truncate">
                  {item.title}
                </h2>
                <p className="text-sm text-slate-500 mt-0.5">
                  {item.price.toLocaleString('ru-RU')} ₽ / шт.
                </p>
              </div>

              {/* Управление количеством (+ / -) */}
              <div className="flex items-center border border-slate-200 rounded-lg overflow-hidden bg-slate-50">
                <button
                  type="button"
                  onClick={() => updateQuantity(item.id, -1)}
                  className="p-1.5 hover:bg-slate-200 text-slate-600 transition-colors"
                  aria-label="Уменьшить"
                >
                  <Minus className="w-4 h-4" />
                </button>
                <span className="w-8 text-center text-sm font-semibold text-slate-800">
                  {item.quantity}
                </span>
                <button
                  type="button"
                  onClick={() => updateQuantity(item.id, 1)}
                  className="p-1.5 hover:bg-slate-200 text-slate-600 transition-colors"
                  aria-label="Увеличить"
                >
                  <Plus className="w-4 h-4" />
                </button>
              </div>

              {/* Сумма по позиции */}
              <div className="w-24 text-right">
                <p className="text-base font-bold text-slate-900">
                  {(item.price * item.quantity).toLocaleString('ru-RU')} ₽
                </p>
              </div>

              {/* Кнопка удаления */}
              <button
                type="button"
                onClick={() => removeFromCart(item.id)}
                className="text-slate-400 hover:text-rose-600 p-2 transition-colors"
                aria-label="Удалить"
              >
                <Trash2 className="w-5 h-5" />
              </button>
            </div>
          ))}
        </div>

        {/* Сводка заказа (1 колонка) */}
        <aside className="lg:col-span-1">
          <div className="bg-white p-6 rounded-xl border border-slate-200 shadow-sm sticky top-24 space-y-5">
            <h2 className="text-lg font-bold text-slate-900">Сумма заказа</h2>

            <div className="space-y-3 border-b border-slate-100 pb-4 text-sm">
              <div className="flex justify-between text-slate-600">
                <span>Товары ({items.reduce((acc, i) => acc + i.quantity, 0)} шт.)</span>
                <span>{rawTotalPrice.toLocaleString('ru-RU')} ₽</span>
              </div>

              {discountPercent > 0 && (
                <div className="flex justify-between text-emerald-600 font-medium">
                  <span>Скидка ({discountPercent}%)</span>
                  <span>−{discountAmount.toLocaleString('ru-RU')} ₽</span>
                </div>
              )}

              <div className="flex justify-between text-slate-600">
                <span>Доставка</span>
                <span className="text-emerald-600 font-medium">Бесплатно</span>
              </div>
            </div>

            {/* Итоговая сумма */}
            <div className="flex justify-between items-baseline">
              <span className="text-base font-medium text-slate-900">Итого к оплате</span>
              <div className="text-right">
                {discountPercent > 0 && (
                  <span className="block text-sm text-slate-400 line-through">
                    {rawTotalPrice.toLocaleString('ru-RU')} ₽
                  </span>
                )}
                <span className="text-2xl font-extrabold text-slate-900">
                  {totalPrice.toLocaleString('ru-RU')} ₽
                </span>
              </div>
            </div>

            {/* Компактный промокод под итоговой ценой */}
            <div className="pt-1">
              <PromoCodeInput />
            </div>

            {/* Кнопка перехода к чекауту */}
            <button
              type="button"
              onClick={() => alert(`Оформление заказа. Промокод: ${promoCode || 'нет'}`)}
              className="w-full bg-emerald-600 text-white py-3.5 px-4 rounded-xl font-bold hover:bg-emerald-500 active:scale-98 transition-all shadow-sm"
            >
              Перейти к оформлению
            </button>

            <Link
              href="/"
              className="mt-4 block text-center text-sm font-medium text-slate-500 hover:text-slate-800 transition-colors"
            >
              Продолжить покупки
            </Link>
          </div>
        </aside>
      </div>
    </div>
  );
}