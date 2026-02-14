import { useQuery } from "@tanstack/react-query";
import { getProductList, getProduct } from "@/api/product";

export function useProducts(page: number, items = 9) {
  return useQuery({
    queryKey: ["products", page, items],
    queryFn: () => getProductList({ page, items }),
  });
}

export function useProduct(id: number) {
  return useQuery({
    queryKey: ["product", id],
    queryFn: () => getProduct(id),
    enabled: id > 0,
  });
}
