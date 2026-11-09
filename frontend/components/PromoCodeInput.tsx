'use client';

import { useState } from 'react';
import { Tag, X, ChevronDown } from 'lucide-react';
import { useCartStore } from '@/store/useCartStore';

export default function PromoCodeInput() {
  const { promoCode, discountPercent, applyPromo, removePromo } = useCartStore();

  const [isOpen, setIsOpen] = useState(false);
  const [inputCode, setInputCode] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const handleApplyPromo = async (e: React.FormEvent) => {
    e.preventDefault();
    const cleanCode = inputCode.trim();
    if (!cleanCode) return;

    setIsLoading(true);
    setErrorMessage('');

    try {
      const response = await fetch('/api/v1/promocodes/apply', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ code: cleanCode }),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.detail || 'Неверный промокод');
      }

      applyPromo(cleanCode, data.discount_percent);
      setInputCode('');
      setIsOpen(false);
    } catch (err: any) {
      setErrorMessage(err.message || 'Ошибка применения');
    } finally {
      setIsLoading(false);
    }
  };

  // Если промокод уже применен — неброский бейдж со скидкой и кнопкой сброса
  if (promoCode) {
    return (
      <div className="flex items-center justify-between py-1.5 px-2 bg-slate-100/80 rounded-md text-xs text-slate-600">
        <div className="flex items-center gap-1.5 font-mono uppercase font-medium">
          <Tag className="w-3.5 h-3.5 text-emerald-600" />
          <span>{promoCode}</span>
          <span className="text-emerald-600 font-semibold">(-{discountPercent}%)</span>
        </div>
        <button
          type="button"
          onClick={removePromo}
          className="text-slate-400 hover:text-rose-600 transition-colors p-0.5"
          title="Удалить промокод"
        >
          <X className="w-3.5 h-3.5" />
        </button>
      </div>
    );
  }

  // Если не раскрыт — компактная строчка-ссылка
  if (!isOpen) {
    return (
      <button
        type="button"
        onClick={() => setIsOpen(true)}
        className="text-xs text-slate-500 hover:text-slate-800 transition-colors inline-flex items-center gap-1 font-medium"
      >
        <span>У вас есть промокод?</span>
        <ChevronDown className="w-3 h-3 text-slate-400" />
      </button>
    );
  }

  // Раскрытая минималистичная форма ввода
  return (
    <div className="space-y-1.5">
      <form onSubmit={handleApplyPromo} className="flex gap-1.5">
        <input
          type="text"
          placeholder="ПРОМОКОД"
          value={inputCode}
          onChange={(e) => setInputCode(e.target.value)}
          autoFocus
          className="flex-1 min-w-0 px-2.5 py-1.5 text-xs uppercase bg-white border border-slate-200 rounded-md focus:outline-none focus:border-slate-400 font-mono"
        />
        <button
          type="submit"
          disabled={isLoading || !inputCode.trim()}
          className="px-3 py-1.5 bg-slate-800 text-white rounded-md text-xs font-medium hover:bg-slate-700 disabled:opacity-40 transition-colors shrink-0"
        >
          {isLoading ? '...' : 'ОК'}
        </button>
        <button
          type="button"
          onClick={() => {
            setIsOpen(false);
            setErrorMessage('');
          }}
          className="p-1.5 text-slate-400 hover:text-slate-600 transition-colors"
          title="Скрыть"
        >
          <X className="w-3.5 h-3.5" />
        </button>
      </form>
      {errorMessage && (
        <p className="text-[11px] text-rose-600 font-medium px-0.5">{errorMessage}</p>
      )}
    </div>
  );
}