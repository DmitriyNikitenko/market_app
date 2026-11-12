from fastapi import APIRouter, Depends
from sqlalchemy.ext.asyncio import AsyncSession

from app.core.database import get_db
from app.schemas.promocode import ApplyPromoRequest, ApplyPromoResponse
from app.services.promo_service import validate_and_apply_promo

router = APIRouter(prefix="/promocodes", tags=["promocodes"])


@router.post("/apply", response_model=ApplyPromoResponse)
async def apply_code(
    payload: ApplyPromoRequest,
    db: AsyncSession = Depends(get_db),
):
    return await validate_and_apply_promo(db, payload)