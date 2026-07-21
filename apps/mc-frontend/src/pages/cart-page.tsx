import { useCartStore } from "@/stores/cart-store";
import { CartEmpty } from "@/features/cart/cart-empty";
import { CartItemList } from "@/features/cart/cart-item-list";
import { CartSummary } from "@/features/cart/cart-summary";

export function CartPage() {
  const cart = useCartStore((s) => s.cart);

  return (
    <div className="mx-auto max-w-4xl px-4 py-8">
      <h1 className="mb-8 text-2xl font-bold">Shopping Cart</h1>

      {cart.length === 0 ? (
        <CartEmpty />
      ) : (
        <div className="flex flex-col gap-8 md:flex-row">
          <div className="flex-1">
            <CartItemList items={cart} />
          </div>
          <div className="w-full md:w-80">
            <CartSummary items={cart} />
          </div>
        </div>
      )}
    </div>
  );
}
