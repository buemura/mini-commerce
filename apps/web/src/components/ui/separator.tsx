import type { HTMLAttributes } from "react";
import { cn } from "@/lib/utils";

export function Separator({
  className,
  ...props
}: HTMLAttributes<HTMLHRElement>) {
  return (
    <hr
      className={cn("border-zinc-200 dark:border-zinc-800", className)}
      {...props}
    />
  );
}
