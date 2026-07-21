import { SignupForm } from "@/features/auth/signup-form";

export function SignupPage() {
  return (
    <div className="mx-auto mt-20 w-full max-w-sm px-4">
      <div className="mb-6 text-center">
        <h1 className="text-2xl font-bold">Sign Up</h1>
        <p className="mt-1 text-sm text-zinc-500">
          Create an account to start shopping
        </p>
      </div>
      <SignupForm />
    </div>
  );
}
