import { useState } from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { QueryClientProvider } from "@tanstack/react-query";
import { createQueryClient } from "@/lib/query-client";
import { Header } from "@/components/layout/header";
import { ToastProvider } from "@/components/layout/toast-provider";
import { ProtectedRoute } from "@/components/layout/protected-route";
import { HomePage } from "@/pages/home-page";
import { SigninPage } from "@/pages/signin-page";
import { SignupPage } from "@/pages/signup-page";
import { ProductsPage } from "@/pages/products-page";
import { ProductDetailPage } from "@/pages/product-detail-page";
import { CartPage } from "@/pages/cart-page";
import { CheckoutPage } from "@/pages/checkout-page";
import { OrdersPage } from "@/pages/orders-page";
import { OrderDetailPage } from "@/pages/order-detail-page";
import { NotFoundPage } from "@/pages/not-found-page";

export function App() {
  const [queryClient] = useState(() => createQueryClient());

  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <div className="min-h-screen bg-white text-zinc-900 dark:bg-zinc-950 dark:text-zinc-100">
          <Header />
          <main>
            <Routes>
              <Route path="/" element={<HomePage />} />
              <Route path="/auth/signin" element={<SigninPage />} />
              <Route path="/auth/signup" element={<SignupPage />} />
              <Route path="/products" element={<ProductsPage />} />
              <Route path="/products/:id" element={<ProductDetailPage />} />
              <Route path="/cart" element={<CartPage />} />
              <Route
                path="/checkout"
                element={
                  <ProtectedRoute>
                    <CheckoutPage />
                  </ProtectedRoute>
                }
              />
              <Route
                path="/orders"
                element={
                  <ProtectedRoute>
                    <OrdersPage />
                  </ProtectedRoute>
                }
              />
              <Route path="/orders/:id" element={<OrderDetailPage />} />
              <Route path="*" element={<NotFoundPage />} />
            </Routes>
          </main>
        </div>
        <ToastProvider />
      </BrowserRouter>
    </QueryClientProvider>
  );
}
