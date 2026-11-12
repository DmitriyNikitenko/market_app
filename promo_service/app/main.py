from fastapi import FastAPI
from app.api.v1.promocodes import router as promo_router
from app.core.config import settings

app = FastAPI(title=settings.PROJECT_NAME)

app.include_router(promo_router, prefix=settings.API_V1_STR)


@app.get("/health")
async def health_check():
    return {"status": "ok", "service": "promo_service"}