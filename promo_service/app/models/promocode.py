from datetime import datetime
from decimal import Decimal
import enum

from sqlalchemy import Boolean, DateTime, Enum, Numeric, String
from sqlalchemy.orm import Mapped, mapped_column

from app.core.database import Base


class DiscountType(str, enum.Enum):
    PERCENT = "percent"
    FIXED = "fixed"


class PromoCode(Base):
    __tablename__ = "promocodes"

    id: Mapped[int] = mapped_column(primary_key=True, index=True)
    code: Mapped[str] = mapped_column(String(50), unique=True, index=True, nullable=False)
    discount_type: Mapped[DiscountType] = mapped_column(Enum(DiscountType), nullable=False)
    discount_value: Mapped[Decimal] = mapped_column(Numeric(10, 2), nullable=False)
    min_order_amount: Mapped[Decimal] = mapped_column(Numeric(10, 2), default=Decimal("0.00"))
    max_discount_amount: Mapped[Decimal | None] = mapped_column(Numeric(10, 2), nullable=True)
    is_active: Mapped[bool] = mapped_column(Boolean, default=True)
    valid_until: Mapped[datetime | None] = mapped_column(DateTime(timezone=True), nullable=True)