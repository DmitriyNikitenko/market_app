'use client';

import { useEffect, useState } from 'react';
import { useCartStore } from '@/store/useCartStore';
import { getProducts, Product } from '@/services/products';

export default function HomePage() {
  const addToCart = useCartStore((state) => state.addToCart);

  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const loadProducts = async () => {
    try {
      setLoading(true);
      setError(null);
      const data = await getProducts();
      setProducts(data);
    } catch (err) {
      setError('Не удалось загрузить каталог');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadProducts();
  }, []);

  return (
    <div className="py-10 px-4 sm:px-6 lg:px-8">
      <div className="max-w-6xl mx-auto">
        <header className="mb-8 border-b border-slate-200 pb-5">
          <h1 className="text-3xl font-extrabold text-slate-900 tracking-tight">Каталог товаров</h1>
          <p className="mt-2 text-sm text-slate-500">Выберите товары, чтобы добавить их в корзину</p>
        </header>

        {loading && (
          <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
            {[1, 2, 3, 4].map((n) => (
              <div key={n} className="h-80 rounded-xl bg-slate-100 animate-pulse border border-slate-200" />
            ))}
          </div>
        )}

        {!loading && error && (
          <div className="rounded-xl border border-transparent p-6 text-center">
            <p className="text-sm font-medium text-red-700">{error}</p>
            <button
              type="button"
              onClick={loadProducts}
              className="mt-3 rounded-lg bg-red-600 px-4 py-2 text-sm font-semibold text-white hover:bg-red-500 cursor-pointer"
            >
              Повторить попытку
            </button>
          </div>
        )}

        {!loading && !error && products.length === 0 && (
          <p className="text-center text-sm text-slate-500 py-12">Список товаров пока пуст</p>
        )}

        {!loading && !error && products.length > 0 && (
          <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
            {products.map((product) => (
              <article key={product.id} className="flex flex-col bg-white rounded-xl shadow-sm border border-slate-200 overflow-hidden">
                <div className="aspect-square w-full bg-slate-100 overflow-hidden">
                  <img src={product.image} alt={product.title} className="w-full h-full object-cover object-center" />
                </div>
                <div className="p-4 flex flex-col flex-1 justify-between">
                  <div>
                    <span className="text-xs font-semibold uppercase text-slate-400">{product.category}</span>
                    <h2 className="mt-1 text-base font-medium text-slate-900 line-clamp-2">{product.title}</h2>
                  </div>
                  <div className="mt-4 pt-3 border-t border-slate-100 flex items-center justify-between">
                    <span className="text-lg font-bold text-slate-900">{product.price.toLocaleString('ru-RU')} ₽</span>
                    <button
                      type="button"
                      onClick={() => addToCart(product)}
                      className="rounded-lg bg-emerald-600 px-3.5 py-2 text-sm font-semibold text-white shadow-sm hover:bg-emerald-500 active:scale-95 transition-all cursor-pointer"
                    >
                      В корзину
                    </button>
                  </div>
                </div>
              </article>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}