from decimal import Decimal
from typing import Literal

from pydantic import BaseModel, Field


class CartItem(BaseModel):
    product_id: int
    price: Decimal = Field(gt=0)
    quantity: int = Field(gt=0)


class ApplyPromoRequest(BaseModel):
    code: str = Field(min_length=1, max_length=50)
    cart_total: Decimal = Field(ge=0)
    items: list[CartItem] = Field(default_factory=list)


class ApplyPromoResponse(BaseModel):
    is_valid: bool
    code: str
    discount_type: Literal["percent", "fixed"] | None = None
    discount_amount: Decimal = Decimal("0.00")
    final_total: Decimal
    message: str