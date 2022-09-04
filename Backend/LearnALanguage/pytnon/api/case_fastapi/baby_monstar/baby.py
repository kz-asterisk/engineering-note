from typing import Optional

from pydantic import BaseModel, Field


class Baby(BaseModel):
    name: str = Field(...)
    birth_month: int = Field(...)
    phys_sex: Optional[str]

    def cry(self):
        return "Gyaaaaa!!!"

    def poop(self):
        return "💩💩💩"

    def milk(self):
        return "🍼🍼🍼"


class BabyFood(BaseModel):
    name: str = Field(...)
