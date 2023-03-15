
from fastapi import Body, FastAPI, status, HTTPException
from fastapi.responses import JSONResponse
from enum import Enum
from typing import Union
from pydantic import BaseModel

from uuid import uuid4


def randomUUID4() -> str:
    return str(uuid4())


class Vote(BaseModel):
    id_: str = str(uuid4())
    question: str
    answers: str


class Pool(BaseModel):
    id_: str = str(uuid4())
    name: str = "new pool"
    votes: list[Vote] = list()


apiDatabase: list[Pool] = list()

app = FastAPI()

@app.get("/")
async def root():
    return {"message": "Hello World"}


@app.get("/pool")
async def getPools():
    return {"elements": apiDatabase}


@app.post("/pool")
async def postPool(pool_name: str or None = None):
    newPool: Pool = Pool()
    if pool_name is not None: newPool.name = pool_name
    apiDatabase.append(newPool)
    return {"pool_id": newPool.id_}


@app.get("/pool/{pool_id}")
async def getPoolByID(pool_id: str):
    try:
        return [x for x in apiDatabase if x.id_ == pool_id][0]

    except:
        raise HTTPException(status_code=404, detail="Element not found")

@app.put("/pool/{pool_id}")
async def putPool(pool_id: str, name_: str):
    try:
        [x for x in apiDatabase if x.id_ == pool_id][0].name = name_
        return HTTPException(status_code=200)

    except:
        raise HTTPException(status_code=404, detail="Element not found")

@app.delete("/pool/{pool_id}")
async def deletePoolByID(pool_id: str):
    apiDatabase.pop([idx for idx, elem in enumerate(
        apiDatabase) if elem.id_ == pool_id][0])
    return HTTPException(status_code=204)

@app.get("/pool/{pool_id}/vote")
async def getPoolVote(pool_id: str):
    try:
        return {"elements": [x for x in apiDatabase if x.id_ == pool_id][0].votes}

    except:
        raise HTTPException(status_code=404, detail="Element not found")

# @app.post("/pool/{pool_id}/vote")   

# @app.get("/pool/{pool_id}/vote{vote_id}")
# @app.put("/pool/{pool_id}/vote{vote_id}")
# @app.delete("/pool/{pool_id}/vote{vote_id}")
