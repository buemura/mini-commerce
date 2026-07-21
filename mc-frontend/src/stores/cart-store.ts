import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";
import type { Product } from "@/types/product";

interface CartState {
  cart: Product[];
  addItem: (item: Product) => void;
  removeItem: (id: number) => void;
  clearCart: () => void;
}

export const useCartStore = create<CartState>()(
  devtools(
    persist(
      (set) => ({
        cart: [],
        addItem: (item) =>
          set((state) => {
            const existing = state.cart.find((p) => p.id === item.id);
            if (existing) {
              return {
                cart: state.cart.map((p) =>
                  p.id === item.id
                    ? { ...p, quantity: p.quantity + item.quantity }
                    : p,
                ),
              };
            }
            return { cart: [...state.cart, item] };
          }),
        removeItem: (id) =>
          set((state) => ({ cart: state.cart.filter((p) => p.id !== id) })),
        clearCart: () => set({ cart: [] }),
      }),
      { name: "shopping-cart" },
    ),
  ),
);
