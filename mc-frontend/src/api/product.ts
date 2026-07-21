import { apiClient } from "@/lib/api-client";
import type {
  GetProductListParams,
  GetProductListResponse,
  Product,
} from "@/types/product";

export function getProductList(
  params: GetProductListParams,
): Promise<GetProductListResponse> {
  return apiClient.get<GetProductListResponse>(
    `/api/products?page=${params.page}&items=${params.items}`,
  );
}

export function getProduct(id: number): Promise<Product> {
  return apiClient.get<Product>(`/api/products/${id}`);
}
