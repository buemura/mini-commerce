import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";
import type { Product } from "@/types/product";

interface CheckoutState {
  checkout: Product[];
  initCheckout: (products: Product[]) => void;
  clearCheckout: () => void;
}

export const useCheckoutStore = create<CheckoutState>()(
  devtools(
    persist(
      (set) => ({
        checkout: [],
        initCheckout: (products) => set({ checkout: products }),
        clearCheckout: () => set({ checkout: [] }),
      }),
      { name: "checkout" },
    ),
  ),
);
