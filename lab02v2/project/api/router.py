from fastapi import APIRouter, Response, Request, Form
from fastapi import HTTPException
from fastapi.templating import Jinja2Templates
from fastapi.responses import RedirectResponse

import logging
# from .models import OrderData, OrderList, OrderListAll, OrderInput, PutResources
from pydantic import ValidationError

from . import weather as WeatherAPI

router = APIRouter()
templates = Jinja2Templates(directory="templates/")

@router.get('/')
async def read_form():
    return RedirectResponse("localhost:8000/form")

@router.get("/form")
async def form_post(request: Request):
    location = "London"
    days = 3
    return templates.TemplateResponse('form.html', context={'request': request, 'location': location, 'days': days})

@router.post("/form")
async def form_post(request: Request, location: str = Form(...), days: int = Form(...)):
    return templates.TemplateResponse('form.html', context={'request': request, 'location': location, 'days': days})
