import os
import random
import sys

from fastapi import FastAPI, APIRouter, Depends, HTTPException
from starlette import status

sys.path.append(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

from config import settings
from models import RequestModel

LOCATIONS = {
    "1": "Living Room",
    "2": "Bedroom",
    "3": "Kitchen",
}

SENSOR_ID = {
    "Living Room" : "1",
    "Bedroom" : "2",
    "Kitchen": "3",
}

router = APIRouter(prefix="", tags=["system"])

@router.get("/temperature")
async def get_random_temperature(data: RequestModel = Depends()):
    temperature = random.randint(settings.min_temperature, settings.max_temperature)
    location, sensor_id = data.location, data.sensorID

    if not data.sensorID and not data.location:
        raise HTTPException(status_code=status.HTTP_404_NOT_FOUND, detail="sensorID или location должны быть определены")
    if data.location:
        sensor_id = SENSOR_ID.get(data.location, "0")
    else:
        location = LOCATIONS.get(data.sensorID, "Unknown")

    return {"temperature": temperature, "location": location, "sensor_id": sensor_id}


app = FastAPI(
    title=settings.name,
    docs_url=settings.docs_url,
)
app.include_router(router)
