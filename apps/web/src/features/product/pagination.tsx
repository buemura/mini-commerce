import { Link, useLocation } from "react-router-dom";
import { ChevronLeft, ChevronRight } from "lucide-react";
import type { PaginationMeta } from "@/types/common";
import { Button } from "@/components/ui/button";

interface PaginationProps {
  meta: PaginationMeta;
}

export function Pagination({ meta }: PaginationProps) {
  const { pathname } = useLocation();

  if (meta.total_pages <= 1) return null;

  const pages = Array.from({ length: meta.total_pages }, (_, i) => i + 1);

  return (
    <nav className="mt-8 flex items-center justify-center gap-1">
      {meta.page > 1 && (
        <Link to={`${pathname}?page=${meta.page - 1}`}>
          <Button variant="ghost" size="icon">
            <ChevronLeft className="h-4 w-4" />
          </Button>
        </Link>
      )}

      {pages.map((page) => (
        <Link key={page} to={`${pathname}?page=${page}`}>
          <Button
            variant={page === meta.page ? "primary" : "ghost"}
            size="icon"
          >
            {page}
          </Button>
        </Link>
      ))}

      {meta.page < meta.total_pages && (
        <Link to={`${pathname}?page=${meta.page + 1}`}>
          <Button variant="ghost" size="icon">
            <ChevronRight className="h-4 w-4" />
          </Button>
        </Link>
      )}
    </nav>
  );
}
