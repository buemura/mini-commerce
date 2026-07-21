import { Link } from "react-router-dom";
import type { Product } from "@/types/product";
import { formatPrice } from "@/utils/currency-formatter";
import { Card } from "@/components/ui/card";

interface ProductCardProps {
  product: Product;
}

export function ProductCard({ product }: ProductCardProps) {
  return (
    <Link to={`/products/${product.id}`}>
      <Card className="overflow-hidden transition-shadow hover:shadow-md">
        <div className="aspect-square overflow-hidden bg-zinc-100 dark:bg-zinc-800">
          <img
            src={product.image_url}
            alt={product.name}
            className="h-full w-full object-cover transition-transform hover:scale-105"
          />
        </div>
        <div className="p-4">
          <h3 className="truncate font-medium">{product.name}</h3>
          <p className="mt-1 text-lg font-semibold">
            {formatPrice(product.price)}
          </p>
        </div>
      </Card>
    </Link>
  );
}
