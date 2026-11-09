'use client';

import { useCartStore } from '@/store/useCartStore';

interface Product {
  id: number;
  title: string;
  category: string;
  price: number;
  image: string;
}

const MOCK_PRODUCTS: Product[] = [
  {
    id: 1,
    title: 'Кофемашина эспрессо с капучинатором',
    category: 'Техника для кухни',
    price: 34990,
    image: 'https://images.unsplash.com/photo-1517668808822-9ebb02f2a0e6?w=600&auto=format&fit=crop&q=60',
  },
  {
    id: 2,
    title: 'Зерновой кофе Arabica Blend (1 кг)',
    category: 'Продукты',
    price: 1890,
    image: 'https://images.unsplash.com/photo-1559056199-641a0ac8b55e?w=600&auto=format&fit=crop&q=60',
  },
  {
    id: 3,
    title: 'Сироп «Французская ваниль» (1 л)',
    category: 'Ингредиенты',
    price: 750,
    image: 'https://images.unsplash.com/photo-1546069901-ba9599a7e63c?w=600&auto=format&fit=crop&q=60',
  },
  {
    id: 4,
    title: 'Набор стеклянных бокалов с двойным дном',
    category: 'Посуда',
    price: 1450,
    image: 'https://images.unsplash.com/photo-1514432324607-a09d9b4aefdd?w=600&auto=format&fit=crop&q=60',
  },
];

export default function HomePage() {
  const addToCart = useCartStore((state) => state.addToCart);

  return (
    <div className="py-10 px-4 sm:px-6 lg:px-8">
      <div className="max-w-6xl mx-auto">
        <header className="mb-8 border-b border-slate-200 pb-5">
          <h1 className="text-3xl font-extrabold text-slate-900 tracking-tight">Каталог товаров</h1>
          <p className="mt-2 text-sm text-slate-500">Нажмите на кнопку, чтобы проверить работу корзины</p>
        </header>

        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
          {MOCK_PRODUCTS.map((product) => (
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
      </div>
    </div>
  );
}