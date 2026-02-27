from uuid import UUID

from fastapi import Query
from pydantic import BaseModel


class RequestModel(BaseModel):
    location: str | None = None
    sensorID: str | None = None

    @classmethod
    def as_query(
            cls,
            sensorID: str | None = Query(..., description="Идентификатор названия комнаты"),
            location: str | None = Query(..., description="Название комнаты"),
    ):
        return cls(location=location, owner_id=owner_id, sensorID=sensorID)
