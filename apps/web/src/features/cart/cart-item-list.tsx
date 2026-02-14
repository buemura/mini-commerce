import { X } from "lucide-react";
import { toast } from "sonner";
import type { Product } from "@/types/product";
import { formatPrice } from "@/utils/currency-formatter";
import { useCartStore } from "@/stores/cart-store";
import { Button } from "@/components/ui/button";

interface CartItemListProps {
  items: Product[];
}

export function CartItemList({ items }: CartItemListProps) {
  const removeItem = useCartStore((s) => s.removeItem);

  const handleRemove = (id: number) => {
    removeItem(id);
    toast.success("Item removed from cart");
  };

  return (
    <div className="flex flex-col gap-4">
      {items.map((item) => (
        <div
          key={item.id}
          className="flex items-center gap-4 rounded-lg border border-zinc-200 p-4 dark:border-zinc-800"
        >
          <div className="h-20 w-20 shrink-0 overflow-hidden rounded-md bg-zinc-100 dark:bg-zinc-800">
            <img
              src={item.image_url}
              alt={item.name}
              className="h-full w-full object-cover"
            />
          </div>

          <div className="flex-1">
            <h3 className="font-medium">{item.name}</h3>
            <p className="text-sm text-zinc-500">Qty: {item.quantity}</p>
            <p className="font-semibold">
              {formatPrice(item.price * item.quantity)}
            </p>
          </div>

          <Button
            variant="ghost"
            size="icon"
            onClick={() => handleRemove(item.id)}
          >
            <X className="h-4 w-4" />
          </Button>
        </div>
      ))}
    </div>
  );
}
