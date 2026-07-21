import type { HTMLAttributes } from "react";
import { cn } from "@/lib/utils";

const variants = {
  default: "bg-zinc-100 text-zinc-800 dark:bg-zinc-800 dark:text-zinc-100",
  success: "bg-emerald-100 text-emerald-800 dark:bg-emerald-900 dark:text-emerald-100",
  warning: "bg-amber-100 text-amber-800 dark:bg-amber-900 dark:text-amber-100",
};

interface BadgeProps extends HTMLAttributes<HTMLSpanElement> {
  variant?: keyof typeof variants;
}

export function Badge({ className, variant = "default", ...props }: BadgeProps) {
  return (
    <span
      className={cn(
        "inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium",
        variants[variant],
        className,
      )}
      {...props}
    />
  );
}
