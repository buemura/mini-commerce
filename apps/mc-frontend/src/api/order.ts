import { apiClient } from "@/lib/api-client";
import type { CreateOrderRequest, GetOrdersResponse, Order } from "@/types/order";

export function createOrder(data: CreateOrderRequest): Promise<Order> {
  return apiClient.post<Order>("/api/orders", data);
}

export function getOrder(id: string): Promise<Order> {
  return apiClient.get<Order>(`/api/orders/${id}`);
}

export function getOrders(params: {
  page: number;
  items: number;
}): Promise<GetOrdersResponse> {
  return apiClient.get<GetOrdersResponse>(
    `/api/orders?page=${params.page}&items=${params.items}`,
  );
}
