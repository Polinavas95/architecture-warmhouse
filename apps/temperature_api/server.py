import os
import random
import sys

from fastapi import FastAPI, APIRouter, Depends

sys.path.append(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

from config import settings
from models import RequestModel


router = APIRouter(prefix="", tags=["system"])

@router.get("/temperature")
async def get_random_temperature(data: RequestModel = Depends()):
    temperature = random.randint(settings.min_temperature, settings.max_temperature)
    return {"temperature": temperature, "location": data.location}


app = FastAPI(
    title=settings.name,
    docs_url=settings.docs_url,
)
app.include_router(router)
print("API_PORT", settings.port)
