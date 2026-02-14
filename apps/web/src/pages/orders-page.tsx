import { useSearchParams } from "react-router-dom";
import { ClipboardList } from "lucide-react";
import { useOrders } from "@/hooks/use-orders";
import { OrderList } from "@/features/order/order-list";
import { Pagination } from "@/features/product/pagination";
import { Skeleton } from "@/components/ui/skeleton";

export function OrdersPage() {
  const [searchParams] = useSearchParams();
  const page = Number(searchParams.get("page")) || 1;
  const { data, isLoading } = useOrders(page);

  return (
    <div className="mx-auto max-w-4xl px-4 py-8">
      <h1 className="mb-8 text-2xl font-bold">My Orders</h1>

      {isLoading ? (
        <div className="flex flex-col gap-4">
          {Array.from({ length: 5 }).map((_, i) => (
            <Skeleton key={i} className="h-20 w-full rounded-lg" />
          ))}
        </div>
      ) : data?.order_list?.length ? (
        <>
          <OrderList orders={data.order_list} />
          <Pagination meta={data.meta} />
        </>
      ) : (
        <div className="flex flex-col items-center gap-4 py-16">
          <ClipboardList className="h-16 w-16 text-zinc-300 dark:text-zinc-600" />
          <p className="text-zinc-500">No orders found.</p>
        </div>
      )}
    </div>
  );
}
