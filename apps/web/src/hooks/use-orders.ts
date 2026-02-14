import { useQuery } from "@tanstack/react-query";
import { getOrders, getOrder } from "@/api/order";

export function useOrders(page: number, items = 10) {
  return useQuery({
    queryKey: ["orders", page, items],
    queryFn: () => getOrders({ page, items }),
  });
}

export function useOrder(id: string) {
  return useQuery({
    queryKey: ["order", id],
    queryFn: () => getOrder(id),
    enabled: !!id,
  });
}
