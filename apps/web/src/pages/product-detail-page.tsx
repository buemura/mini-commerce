import { useParams } from "react-router-dom";
import { useProduct } from "@/hooks/use-products";
import { formatPrice } from "@/utils/currency-formatter";
import { ProductForm } from "@/features/product/product-form";
import { Skeleton } from "@/components/ui/skeleton";

export function ProductDetailPage() {
  const { id } = useParams();
  const { data: product, isLoading } = useProduct(Number(id));

  if (isLoading) {
    return (
      <div className="mx-auto flex max-w-6xl flex-col gap-8 px-4 py-8 md:flex-row">
        <div className="flex-1">
          <Skeleton className="h-8 w-3/4" />
          <Skeleton className="mt-2 h-10 w-1/4" />
          <Skeleton className="mt-4 h-5 w-1/3" />
          <Skeleton className="mt-8 h-10 w-48" />
          <Skeleton className="mt-8 h-20 w-full" />
        </div>
        <Skeleton className="aspect-square w-full max-w-md rounded-lg" />
      </div>
    );
  }

  if (!product) {
    return (
      <div className="py-16 text-center">
        <p className="text-zinc-500">Product not found.</p>
      </div>
    );
  }

  return (
    <div className="mx-auto flex max-w-6xl flex-col-reverse gap-8 px-4 py-8 md:flex-row">
      <div className="flex-1">
        <h1 className="text-3xl font-bold">{product.name}</h1>
        <p className="mt-2 text-2xl font-semibold">
          {formatPrice(product.price)}
        </p>
        <p className="mt-1 text-sm text-zinc-500">
          {product.quantity > 0
            ? `${product.quantity} in stock`
            : "Out of stock"}
        </p>

        <div className="mt-6">
          <ProductForm product={product} />
        </div>

        <div className="mt-8">
          <h2 className="text-lg font-semibold">Description</h2>
          <p className="mt-2 text-zinc-600 dark:text-zinc-400">
            {product.description}
          </p>
        </div>
      </div>

      <div className="w-full max-w-md">
        <div className="aspect-square overflow-hidden rounded-lg bg-zinc-100 dark:bg-zinc-800">
          <img
            src={product.image_url}
            alt={product.name}
            className="h-full w-full object-cover"
          />
        </div>
      </div>
    </div>
  );
}
