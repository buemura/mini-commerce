import { useState } from "react";
import {
  formatCardNumber,
  formatExpiry,
  formatCVV,
} from "@/utils/credit-card-formatter";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

export function CreditCardForm() {
  const [cardNumber, setCardNumber] = useState("");
  const [expiry, setExpiry] = useState("");
  const [cvv, setCvv] = useState("");

  return (
    <div className="flex flex-col gap-4">
      <div className="flex flex-col gap-2">
        <Label htmlFor="card-number">Card Number</Label>
        <Input
          id="card-number"
          placeholder="0000 0000 0000 0000"
          value={cardNumber}
          onChange={(e) => setCardNumber(formatCardNumber(e.target.value))}
        />
      </div>

      <div className="flex gap-4">
        <div className="flex flex-1 flex-col gap-2">
          <Label htmlFor="expiry">Expiry</Label>
          <Input
            id="expiry"
            placeholder="MM/YY"
            value={expiry}
            onChange={(e) => setExpiry(formatExpiry(e.target.value))}
          />
        </div>

        <div className="flex w-24 flex-col gap-2">
          <Label htmlFor="cvv">CVV</Label>
          <Input
            id="cvv"
            placeholder="000"
            value={cvv}
            onChange={(e) => setCvv(formatCVV(e.target.value))}
          />
        </div>
      </div>
    </div>
  );
}
