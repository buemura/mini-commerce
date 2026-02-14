import { useNavigate } from "react-router-dom";
import { toast } from "sonner";
import type { Product } from "@/types/product";
import { formatPrice } from "@/utils/currency-formatter";
import { useUserStore } from "@/stores/user-store";
import { useCheckoutStore } from "@/stores/checkout-store";
import { useOrderStore } from "@/stores/order-store";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

interface CartSummaryProps {
  items: Product[];
}

export function CartSummary({ items }: CartSummaryProps) {
  const navigate = useNavigate();
  const user = useUserStore((s) => s.user);
  const initCheckout = useCheckoutStore((s) => s.initCheckout);
  const { setCustomerId, setProductList } = useOrderStore();

  const total = items.reduce((sum, item) => sum + item.price * item.quantity, 0);

  const handleProceed = () => {
    if (!user) {
      toast.error("Please sign in to continue");
      navigate("/auth/signin");
      return;
    }

    initCheckout(items);
    setCustomerId(user.customer.id);
    setProductList(items);
    navigate("/checkout");
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle>Order Summary</CardTitle>
      </CardHeader>
      <CardContent className="flex flex-col gap-4">
        <div className="flex items-center justify-between">
          <span className="text-zinc-500">Total ({items.length} items)</span>
          <span className="text-lg font-semibold">{formatPrice(total)}</span>
        </div>
        <Separator />
        <Button onClick={handleProceed} className="w-full">
          Proceed to Checkout
        </Button>
      </CardContent>
    </Card>
  );
}
