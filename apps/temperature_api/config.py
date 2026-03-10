from typing import Annotated

from pydantic import Field
from pydantic_settings import BaseSettings


class ApiConfig(BaseSettings):
    name: Annotated[str, Field(alias="API_NAME")] = "Temperature API"
    description: Annotated[str, Field(alias="API_DESCRIPTION")] = "Приложение для получения рандомной температуры"
    docs_url: Annotated[str, Field(alias="API_DOCS_URL")] = "/"
    host: Annotated[str, Field(alias="API_HOST")] = "127.0.0.1"
    port: Annotated[int, Field(alias="API_PORT")] = 8080
    min_temperature: Annotated[int, Field(alias="MIN_TEMPERATURE")] = 5
    max_temperature: Annotated[int, Field(alias="MAX_TEMPERATURE")] = 35

settings = ApiConfig()
