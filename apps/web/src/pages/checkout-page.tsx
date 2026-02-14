import { PaymentMethodSelect } from "@/features/checkout/payment-method-select";
import { CheckoutProductReview } from "@/features/checkout/checkout-product-review";
import { CheckoutFinishButton } from "@/features/checkout/checkout-finish-button";

export function CheckoutPage() {
  return (
    <div className="mx-auto max-w-4xl px-4 py-8">
      <h1 className="mb-8 text-2xl font-bold">Checkout</h1>

      <div className="flex flex-col gap-8 md:flex-row">
        <div className="flex-1">
          <PaymentMethodSelect />
        </div>
        <div className="w-full md:w-80">
          <div className="flex flex-col gap-4">
            <CheckoutProductReview />
            <CheckoutFinishButton />
          </div>
        </div>
      </div>
    </div>
  );
}
