from datetime import datetime, timezone
from decimal import Decimal
from sqlalchemy import select
from sqlalchemy.ext.asyncio import AsyncSession

from app.models.promocode import DiscountType, PromoCode
from app.schemas.promocode import ApplyPromoRequest, ApplyPromoResponse


async def validate_and_apply_promo(
    db: AsyncSession, request: ApplyPromoRequest
) -> ApplyPromoResponse:
    cleaned_code = request.code.strip().upper()
    query = select(PromoCode).where(PromoCode.code == cleaned_code)
    result = await db.execute(query)
    promo = result.scalar_one_or_none()

    if not promo or not promo.is_active:
        return ApplyPromoResponse(
            is_valid=False,
            code=request.code,
            final_total=request.cart_total,
            message="Промокод не найден или неактивен",
        )

    if promo.valid_until and promo.valid_until < datetime.now(timezone.utc):
        return ApplyPromoResponse(
            is_valid=False,
            code=request.code,
            final_total=request.cart_total,
            message="Срок действия промокода истек",
        )

    if request.cart_total < promo.min_order_amount:
        return ApplyPromoResponse(
            is_valid=False,
            code=request.code,
            final_total=request.cart_total,
            message=f"Минимальная сумма заказа для применения: {promo.min_order_amount} ₽",
        )

    if promo.discount_type == DiscountType.PERCENT:
        discount = (request.cart_total * promo.discount_value) / Decimal("100")
        if promo.max_discount_amount and discount > promo.max_discount_amount:
            discount = promo.max_discount_amount
    else:
        discount = promo.discount_value

    discount = min(discount, request.cart_total)
    final_total = max(Decimal("0.00"), request.cart_total - discount)

    return ApplyPromoResponse(
        is_valid=True,
        code=promo.code,
        discount_type=promo.discount_type.value,
        discount_amount=discount.quantize(Decimal("0.01")),
        final_total=final_total.quantize(Decimal("0.01")),
        message="Промокод успешно применен",
    )