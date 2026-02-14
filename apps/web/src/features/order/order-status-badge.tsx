import { OrderStatus } from "@/utils/constant-mapper";
import { Badge } from "@/components/ui/badge";

interface OrderStatusBadgeProps {
  status: string;
}

export function OrderStatusBadge({ status }: OrderStatusBadgeProps) {
  const variant = status === "COMPLETED" ? "success" : "warning";
  const label = OrderStatus[status] || status;

  return <Badge variant={variant}>{label}</Badge>;
}
