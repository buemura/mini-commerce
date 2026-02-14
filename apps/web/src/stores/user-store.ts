import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";
import type { CustomerAuth } from "@/types/customer";

interface UserState {
  user: CustomerAuth | null;
  setUser: (auth: CustomerAuth) => void;
  logoutUser: () => void;
}

export const useUserStore = create<UserState>()(
  devtools(
    persist(
      (set) => ({
        user: null,
        setUser: (auth) => set({ user: auth }),
        logoutUser: () => set({ user: null }),
      }),
      { name: "user-storage" },
    ),
  ),
);
