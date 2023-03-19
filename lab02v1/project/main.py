from fastapi import FastAPI
from api.router import router

app = FastAPI(
    version="0.0.1",
    title="mrrys00 random numbers proxy service",
    description="API that connects to "
)
app.include_router(router)