import { Link } from "react-router-dom";
import { ArrowRight } from "lucide-react";
import { Button } from "@/components/ui/button";

export function HomePage() {
  return (
    <div className="flex min-h-[calc(100vh-3.5rem)] flex-col items-center justify-center gap-8 px-4 text-center">
      <div className="max-w-2xl">
        <h1 className="text-4xl font-bold tracking-tight sm:text-5xl lg:text-6xl">
          Shop the Latest Trends
        </h1>
        <p className="mt-4 text-lg text-zinc-500 dark:text-zinc-400">
          Discover our curated collection of premium products. From everyday
          essentials to exclusive finds, we have everything you need.
        </p>
        <div className="mt-8">
          <Link to="/products">
            <Button size="lg">
              View Products
              <ArrowRight className="h-4 w-4" />
            </Button>
          </Link>
        </div>
      </div>
    </div>
  );
}
