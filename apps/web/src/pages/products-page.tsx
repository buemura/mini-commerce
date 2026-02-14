import { useSearchParams } from "react-router-dom";
import { useProducts } from "@/hooks/use-products";
import { ProductGrid } from "@/features/product/product-grid";
import { Pagination } from "@/features/product/pagination";
import { Skeleton } from "@/components/ui/skeleton";

export function ProductsPage() {
  const [searchParams] = useSearchParams();
  const page = Number(searchParams.get("page")) || 1;
  const { data, isLoading } = useProducts(page);

  return (
    <div className="mx-auto max-w-6xl px-4 py-8">
      <h1 className="mb-8 text-2xl font-bold">Products</h1>

      {isLoading ? (
        <div className="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
          {Array.from({ length: 9 }).map((_, i) => (
            <div key={i} className="flex flex-col gap-3">
              <Skeleton className="aspect-square w-full rounded-lg" />
              <Skeleton className="h-5 w-3/4" />
              <Skeleton className="h-6 w-1/4" />
            </div>
          ))}
        </div>
      ) : data?.product_list?.length ? (
        <>
          <ProductGrid products={data.product_list} />
          <Pagination meta={data.meta} />
        </>
      ) : (
        <p className="py-16 text-center text-zinc-500">No products found.</p>
      )}
    </div>
  );
}
