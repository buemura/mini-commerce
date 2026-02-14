import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { toast } from "sonner";
import { loginCustomer, registerCustomer } from "@/api/customer";
import { useUserStore } from "@/stores/user-store";

export function useSignIn() {
  const navigate = useNavigate();
  const setUser = useUserStore((s) => s.setUser);

  return useMutation({
    mutationFn: loginCustomer,
    onSuccess: (data) => {
      setUser(data);
      toast.success("Signed in successfully");
      navigate("/");
    },
    onError: () => {
      toast.error("Failed to sign in. Please check your credentials.");
    },
  });
}

export function useSignUp() {
  const navigate = useNavigate();

  return useMutation({
    mutationFn: registerCustomer,
    onSuccess: () => {
      toast.success("Account created successfully");
      navigate("/auth/signin");
    },
    onError: () => {
      toast.error("Failed to create account. Please try again.");
    },
  });
}
