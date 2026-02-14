import { SigninForm } from "@/features/auth/signin-form";

export function SigninPage() {
  return (
    <div className="mx-auto mt-20 w-full max-w-sm px-4">
      <div className="mb-6 text-center">
        <h1 className="text-2xl font-bold">Sign In</h1>
        <p className="mt-1 text-sm text-zinc-500">
          Enter your credentials to access your account
        </p>
      </div>
      <SigninForm />
    </div>
  );
}
