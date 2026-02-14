import { Link } from "react-router-dom";
import type { Order } from "@/types/order";
import { formatPrice } from "@/utils/currency-formatter";
import { PaymentMethod } from "@/utils/constant-mapper";
import { Card } from "@/components/ui/card";
import { OrderStatusBadge } from "./order-status-badge";

interface OrderListProps {
  orders: Order[];
}

export function OrderList({ orders }: OrderListProps) {
  return (
    <div className="flex flex-col gap-4">
      {orders.map((order) => (
        <Link key={order.id} to={`/orders/${order.id}`}>
          <Card className="flex items-center justify-between p-4 transition-shadow hover:shadow-md">
            <div className="flex flex-col gap-1">
              <p className="text-sm font-medium">
                Order #{order.id.slice(0, 8)}...
              </p>
              <p className="text-xs text-zinc-500">
                {new Date(order.created_at).toLocaleDateString()} &middot;{" "}
                {PaymentMethod[order.payment_method] || order.payment_method}
              </p>
            </div>

            <div className="flex items-center gap-4">
              <OrderStatusBadge status={order.status} />
              <p className="font-semibold">
                {formatPrice(order.total_price)}
              </p>
            </div>
          </Card>
        </Link>
      ))}
    </div>
  );
}
