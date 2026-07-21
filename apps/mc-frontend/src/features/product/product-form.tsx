import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { toast } from "sonner";
import type { Product } from "@/types/product";
import { useUserStore } from "@/stores/user-store";
import { useCartStore } from "@/stores/cart-store";
import { useCheckoutStore } from "@/stores/checkout-store";
import { useOrderStore } from "@/stores/order-store";
import { Button } from "@/components/ui/button";
import { Select } from "@/components/ui/select";
import { Label } from "@/components/ui/label";

interface ProductFormProps {
  product: Product;
}

export function ProductForm({ product }: ProductFormProps) {
  const [quantity, setQuantity] = useState(1);
  const navigate = useNavigate();
  const user = useUserStore((s) => s.user);
  const addItem = useCartStore((s) => s.addItem);
  const initCheckout = useCheckoutStore((s) => s.initCheckout);
  const { setCustomerId, setProductList } = useOrderStore();

  const handleBuy = () => {
    if (!user) {
      toast.error("Please sign in to continue");
      navigate("/auth/signin");
      return;
    }

    const item = { ...product, quantity };
    initCheckout([item]);
    setCustomerId(user.customer.id);
    setProductList([item]);
    navigate("/checkout");
  };

  const handleAddToCart = () => {
    addItem({ ...product, quantity });
    toast.success("Item added to cart");
  };

  return (
    <div className="flex flex-col gap-4">
      <div className="flex flex-col gap-2">
        <Label htmlFor="quantity">Quantity</Label>
        <Select
          id="quantity"
          value={quantity}
          onChange={(e) => setQuantity(Number(e.target.value))}
          className="w-24"
        >
          {[1, 2, 3, 4, 5].map((n) => (
            <option key={n} value={n}>
              {n}
            </option>
          ))}
        </Select>
      </div>

      <div className="flex gap-3">
        <Button onClick={handleBuy}>Buy Now</Button>
        <Button variant="outline" onClick={handleAddToCart}>
          Add to Cart
        </Button>
      </div>
    </div>
  );
}
