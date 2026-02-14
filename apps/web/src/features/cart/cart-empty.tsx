import { Link } from "react-router-dom";
import { ShoppingCart } from "lucide-react";
import { Button } from "@/components/ui/button";

export function CartEmpty() {
  return (
    <div className="flex flex-col items-center gap-4 py-16">
      <ShoppingCart className="h-16 w-16 text-zinc-300 dark:text-zinc-600" />
      <h2 className="text-xl font-semibold">Your cart is empty</h2>
      <p className="text-zinc-500">Browse our products and add items to your cart.</p>
      <Link to="/products">
        <Button>View Products</Button>
      </Link>
    </div>
  );
}
