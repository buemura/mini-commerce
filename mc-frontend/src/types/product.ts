import type { PaginationMeta } from "./common";

export interface Product {
  id: number;
  name: string;
  description: string;
  price: number;
  quantity: number;
  image_url: string;
}

export interface GetProductListParams {
  page: number;
  items: number;
}

export interface GetProductListResponse {
  product_list: Product[];
  meta: PaginationMeta;
}
