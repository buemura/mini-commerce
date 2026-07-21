import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";
import type { Product } from "@/types/product";

interface OrderState {
  order: {
    customer_id: string;
    payment_method: string;
    product_list: Product[];
  };
  setCustomerId: (id: string) => void;
  setPaymentMethod: (paymentMethod: string) => void;
  setProductList: (productList: Product[]) => void;
  clearOrder: () => void;
}

const initialOrder = {
  customer_id: "",
  payment_method: "",
  product_list: [] as Product[],
};

export const useOrderStore = create<OrderState>()(
  devtools(
    persist(
      (set) => ({
        order: { ...initialOrder },
        setCustomerId: (id) =>
          set((state) => ({ order: { ...state.order, customer_id: id } })),
        setPaymentMethod: (paymentMethod) =>
          set((state) => ({
            order: { ...state.order, payment_method: paymentMethod },
          })),
        setProductList: (productList) =>
          set((state) => ({
            order: { ...state.order, product_list: productList },
          })),
        clearOrder: () => set({ order: { ...initialOrder } }),
      }),
      { name: "order" },
    ),
  ),
);
