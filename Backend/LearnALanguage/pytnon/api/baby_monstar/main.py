from fastapi import FastAPI

from baby_monstar import baby

app = FastAPI()


@app.get("/hello")
async def hello():
    baby.cry()
    return {"message": "hello world!"}
