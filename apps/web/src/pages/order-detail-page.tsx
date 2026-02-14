import { useParams } from "react-router-dom";
import { useOrder } from "@/hooks/use-orders";
import { formatPrice } from "@/utils/currency-formatter";
import { PaymentMethod } from "@/utils/constant-mapper";
import { OrderStatusBadge } from "@/features/order/order-status-badge";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";
import { Skeleton } from "@/components/ui/skeleton";

export function OrderDetailPage() {
  const { id } = useParams();
  const { data: order, isLoading } = useOrder(id || "");

  if (isLoading) {
    return (
      <div className="mx-auto max-w-2xl px-4 py-8">
        <Skeleton className="mb-8 h-8 w-48" />
        <Skeleton className="h-64 w-full rounded-lg" />
      </div>
    );
  }

  if (!order) {
    return (
      <div className="py-16 text-center">
        <p className="text-zinc-500">Order not found.</p>
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-2xl px-4 py-8">
      <h1 className="mb-8 text-2xl font-bold">Order Details</h1>

      <Card>
        <CardHeader>
          <div className="flex items-center justify-between">
            <CardTitle>Order #{order.id.slice(0, 8)}...</CardTitle>
            <OrderStatusBadge status={order.status} />
          </div>
        </CardHeader>
        <CardContent className="flex flex-col gap-4">
          <div className="grid grid-cols-2 gap-4 text-sm">
            <div>
              <p className="text-zinc-500">Date</p>
              <p className="font-medium">
                {new Date(order.created_at).toLocaleString()}
              </p>
            </div>
            <div>
              <p className="text-zinc-500">Payment Method</p>
              <p className="font-medium">
                {PaymentMethod[order.payment_method] || order.payment_method}
              </p>
            </div>
          </div>

          <Separator />

          <div>
            <p className="mb-2 text-sm font-medium text-zinc-500">Products</p>
            {order.product_list.map((item) => (
              <div
                key={item.id}
                className="flex items-center justify-between py-2"
              >
                <div>
                  <p className="text-sm">Product #{item.id}</p>
                  <p className="text-xs text-zinc-500">Qty: {item.quantity}</p>
                </div>
                <p className="text-sm font-medium">
                  {formatPrice(item.price * item.quantity)}
                </p>
              </div>
            ))}
          </div>

          <Separator />

          <div className="flex items-center justify-between">
            <span className="font-medium">Total</span>
            <span className="text-lg font-semibold">
              {formatPrice(order.total_price)}
            </span>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
