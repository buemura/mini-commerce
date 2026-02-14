import { Link } from "react-router-dom";
import {
  PackageOpen,
  ShoppingCart,
  ClipboardList,
  LogIn,
  LogOut,
  Sun,
  Moon,
  Monitor,
} from "lucide-react";
import { useUserStore } from "@/stores/user-store";
import { useCartStore } from "@/stores/cart-store";
import { useTheme } from "@/hooks/use-theme";
import { Button } from "@/components/ui/button";
import { Avatar } from "@/components/ui/avatar";

export function Header() {
  const user = useUserStore((s) => s.user);
  const logoutUser = useUserStore((s) => s.logoutUser);
  const cartCount = useCartStore((s) => s.cart.length);
  const { theme, setTheme } = useTheme();

  const cycleTheme = () => {
    const next = theme === "light" ? "dark" : theme === "dark" ? "system" : "light";
    setTheme(next);
  };

  const themeIcon =
    theme === "dark" ? (
      <Moon className="h-4 w-4" />
    ) : theme === "light" ? (
      <Sun className="h-4 w-4" />
    ) : (
      <Monitor className="h-4 w-4" />
    );

  return (
    <header className="sticky top-0 z-50 border-b border-zinc-200 bg-white/80 backdrop-blur dark:border-zinc-800 dark:bg-zinc-950/80">
      <div className="mx-auto flex h-14 max-w-6xl items-center justify-between px-4">
        <Link to="/" className="flex items-center gap-2 font-semibold">
          <PackageOpen className="h-5 w-5" />
          D-Commerce
        </Link>

        <nav className="flex items-center gap-1">
          <Link to="/orders">
            <Button variant="ghost" size="icon" title="Orders">
              <ClipboardList className="h-4 w-4" />
            </Button>
          </Link>

          <Link to="/cart" className="relative">
            <Button variant="ghost" size="icon" title="Cart">
              <ShoppingCart className="h-4 w-4" />
            </Button>
            {cartCount > 0 && (
              <span className="absolute -right-0.5 -top-0.5 flex h-4 w-4 items-center justify-center rounded-full bg-zinc-900 text-[10px] font-medium text-white dark:bg-zinc-100 dark:text-zinc-900">
                {cartCount}
              </span>
            )}
          </Link>

          <Button variant="ghost" size="icon" onClick={cycleTheme} title="Theme">
            {themeIcon}
          </Button>

          {user ? (
            <div className="flex items-center gap-2">
              <Avatar name={user.customer.name} />
              <Button
                variant="ghost"
                size="icon"
                onClick={logoutUser}
                title="Sign out"
              >
                <LogOut className="h-4 w-4" />
              </Button>
            </div>
          ) : (
            <Link to="/auth/signin">
              <Button variant="ghost" size="icon" title="Sign in">
                <LogIn className="h-4 w-4" />
              </Button>
            </Link>
          )}
        </nav>
      </div>
    </header>
  );
}
