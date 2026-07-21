export interface Customer {
  id: string;
  name: string;
  email: string;
}

export interface SignInRequest {
  email: string;
  password: string;
}

export interface SignInResponse {
  access_token: string;
  customer: Customer;
}

export interface SignUpRequest {
  name: string;
  email: string;
  password: string;
}

export interface CustomerAuth extends SignInResponse {}
