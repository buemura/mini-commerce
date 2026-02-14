import { useCheckoutStore } from "@/stores/checkout-store";
import { formatPrice } from "@/utils/currency-formatter";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";

export function CheckoutProductReview() {
  const checkout = useCheckoutStore((s) => s.checkout);
  const total = checkout.reduce(
    (sum, item) => sum + item.price * item.quantity,
    0,
  );

  return (
    <Card>
      <CardHeader>
        <CardTitle>Product Review</CardTitle>
      </CardHeader>
      <CardContent className="flex flex-col gap-4">
        {checkout.map((item) => (
          <div key={item.id} className="flex items-center gap-3">
            <div className="h-12 w-12 shrink-0 overflow-hidden rounded-md bg-zinc-100 dark:bg-zinc-800">
              <img
                src={item.image_url}
                alt={item.name}
                className="h-full w-full object-cover"
              />
            </div>
            <div className="flex-1">
              <p className="text-sm font-medium">{item.name}</p>
              <p className="text-xs text-zinc-500">Qty: {item.quantity}</p>
            </div>
            <p className="text-sm font-semibold">
              {formatPrice(item.price * item.quantity)}
            </p>
          </div>
        ))}

        <Separator />

        <div className="flex items-center justify-between">
          <span className="font-medium">Total</span>
          <span className="text-lg font-semibold">{formatPrice(total)}</span>
        </div>
      </CardContent>
    </Card>
  );
}
