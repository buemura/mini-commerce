import { useUserStore } from "@/stores/user-store";

const BASE_URL = import.meta.env.VITE_API_BASE_URL || "";

async function request<T>(url: string, options: RequestInit = {}): Promise<T> {
  const token = useUserStore.getState().user?.access_token;

  const headers: Record<string, string> = {
    "Content-Type": "application/json",
    ...((options.headers as Record<string, string>) || {}),
  };

  if (token) {
    headers["Authorization"] = `Bearer ${token}`;
  }

  const response = await fetch(`${BASE_URL}${url}`, { ...options, headers });

  if (!response.ok) {
    const message = await response.text().catch(() => "Request failed");
    throw new Error(message);
  }

  const text = await response.text();
  return text ? JSON.parse(text) : ({} as T);
}

export const apiClient = {
  get: <T>(url: string) => request<T>(url),
  post: <T>(url: string, body: unknown) =>
    request<T>(url, { method: "POST", body: JSON.stringify(body) }),
};
