from uuid import UUID

from fastapi import Query
from pydantic import BaseModel


class RequestModel(BaseModel):
    location: str
    owner_id: UUID

    @classmethod
    def as_query(cls,
                 location: str = Query(..., description="Город или локация"),
                 owner_id: UUID = Query(..., description="ID владельца")):
        return cls(location=location, owner_id=owner_id)
