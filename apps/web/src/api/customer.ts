import { apiClient } from "@/lib/api-client";
import type {
  Customer,
  SignInRequest,
  SignInResponse,
  SignUpRequest,
} from "@/types/customer";

export function loginCustomer(data: SignInRequest): Promise<SignInResponse> {
  return apiClient.post<SignInResponse>("/api/auth/signin", data);
}

export function registerCustomer(data: SignUpRequest): Promise<void> {
  return apiClient.post<void>("/api/auth/signup", data);
}

export function getCustomer(): Promise<Customer> {
  return apiClient.get<Customer>("/api/customer");
}
