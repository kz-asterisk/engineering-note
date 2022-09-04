from fastapi import FastAPI

from baby_monstar import baby

app = FastAPI()
my_baby = baby.Baby(name="bebi-chang", birth_month=2, phys_sex="men")


@app.get("/")
async def get_hello():
    return {"message": "ここには何もないよ"}


@app.get("/hello")
async def get_hello():
    return {"message": "Hi! I'm baby monser👶👶👶"}


@app.get("/cry")
async def get_cry():
    print(my_baby.cry())
    return {"message": my_baby.cry()}


@app.get("/poop")
async def get_poop():
    return {"message": my_baby.poop()}


@app.get("/milk")
async def get_milk():
    return {"message": my_baby.milk()}


@app.post("/babyfood/{babyfood_id}")
def uadd_item(babyfood_id: int, baby_food: baby.BabyFood):
    return {"babyfood_id": babyfood_id, "baby_food_name": baby_food.name}
