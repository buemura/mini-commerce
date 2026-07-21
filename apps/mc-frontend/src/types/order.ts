import type { PaginationMeta } from "./common";
import type { Product } from "./product";

export interface CreateOrderRequest {
  customer_id: string;
  product_list: Product[];
  payment_method: string;
}

export interface OrderProduct {
  id: number;
  price: number;
  quantity: number;
}

export interface Order {
  id: string;
  customer_id: string;
  product_list: OrderProduct[];
  total_price: number;
  status: string;
  payment_method: string;
  created_at: string;
  updated_at: string;
}

export interface GetOrdersResponse {
  order_list: Order[];
  meta: PaginationMeta;
}
