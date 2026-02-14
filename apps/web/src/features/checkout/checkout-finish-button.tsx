import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { Loader2 } from "lucide-react";
import { toast } from "sonner";
import { createOrder } from "@/api/order";
import { useOrderStore } from "@/stores/order-store";
import { useCheckoutStore } from "@/stores/checkout-store";
import { useCartStore } from "@/stores/cart-store";
import { Button } from "@/components/ui/button";

export function CheckoutFinishButton() {
  const navigate = useNavigate();
  const order = useOrderStore((s) => s.order);
  const clearOrder = useOrderStore((s) => s.clearOrder);
  const clearCheckout = useCheckoutStore((s) => s.clearCheckout);
  const clearCart = useCartStore((s) => s.clearCart);

  const isDisabled =
    !order.customer_id ||
    !order.payment_method ||
    order.product_list.length === 0;

  const { mutate, isPending } = useMutation({
    mutationFn: createOrder,
    onSuccess: (data) => {
      clearOrder();
      clearCheckout();
      clearCart();
      toast.success("Order placed successfully!");
      navigate(`/orders/${data.id}`);
    },
    onError: () => {
      toast.error("Failed to place order. Please try again.");
    },
  });

  const handleFinish = () => {
    mutate({
      customer_id: order.customer_id,
      product_list: order.product_list,
      payment_method: order.payment_method,
    });
  };

  return (
    <Button
      onClick={handleFinish}
      disabled={isDisabled || isPending}
      className="w-full"
    >
      {isPending && <Loader2 className="h-4 w-4 animate-spin" />}
      Finish Checkout
    </Button>
  );
}
