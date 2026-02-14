import { useOrderStore } from "@/stores/order-store";
import { Select } from "@/components/ui/select";
import { Label } from "@/components/ui/label";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { CreditCardForm } from "./credit-card-form";

export function PaymentMethodSelect() {
  const paymentMethod = useOrderStore((s) => s.order.payment_method);
  const setPaymentMethod = useOrderStore((s) => s.setPaymentMethod);

  return (
    <Card>
      <CardHeader>
        <CardTitle>Payment Information</CardTitle>
      </CardHeader>
      <CardContent className="flex flex-col gap-4">
        <div className="flex flex-col gap-2">
          <Label htmlFor="payment-method">Payment Method</Label>
          <Select
            id="payment-method"
            value={paymentMethod}
            onChange={(e) => setPaymentMethod(e.target.value)}
          >
            <option value="">Select a method</option>
            <option value="PIX">Pix</option>
            <option value="CREDIT_CARD">Credit Card</option>
          </Select>
        </div>

        {paymentMethod === "CREDIT_CARD" && <CreditCardForm />}
      </CardContent>
    </Card>
  );
}
