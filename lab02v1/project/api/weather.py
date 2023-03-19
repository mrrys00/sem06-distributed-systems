from requests import post, Response
import json

from . import constants as config

defaultURL: str = "https://api.m3o.com/v1/weather/"

def getWeatherForecast(days: int, location: str) -> Response:
    # Get weather forecast for specific number of days
    url = f'{defaultURL}/Forecast'

    data: dict = {
        "days": days,
        "location": location
    }
    resp = post(url, json=data, headers=config.weatherToken)
    # printAll(name="PostCompany", response=resp, data=data)
    return resp

def getWeatherForecastNow(location: str) -> Response:
    # Get weather forecast for now
    url = f'{defaultURL}/Forecast'

    data: dict = {
        "location": location
    }
    resp = post(url, json=data, headers=config.weatherToken)
    # printAll(name="PostCompany", response=resp, data=data)
    return resp
