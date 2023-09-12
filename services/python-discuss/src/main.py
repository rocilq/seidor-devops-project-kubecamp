import datetime
from fastapi import FastAPI, Header, Response
from fastapi.responses import JSONResponse
from pymongo import MongoClient
from dotenv import load_dotenv
from pydantic import BaseModel, Field
from typing import List, Optional
import os
import requests


class DiscussionInput(BaseModel):
    lessonId: str
    title: str
    content: str


class Comment(BaseModel):
    title: str
    content: str
    user: Optional[str] = "Anonymous"
    createdAt: datetime.datetime


class DiscussionOutput(BaseModel):
    dicussionsId: str = Field(alias="_id")
    lessonId: str
    comments: List[Comment]


def parse_object_id(discussion):
    discussion["_id"] = str(discussion["_id"])
    return discussion


load_dotenv()

auth_service = os.getenv("AUTH_SERVICE")
host = os.getenv("MONGODB_HOST", "localhost")
port = os.getenv("MONGODB_PORT", 27017)
client = MongoClient(host, int(port))
db = os.getenv("MONGODB_DB", "kubecampDiscuss")
db = client[db]
discussions = db.discussions

app = FastAPI()


@app.get("/")
def index():
    return {"message": "Server is up and running"}


@app.get("/discussions/{lessonId}", response_model=DiscussionOutput)
def get_discussion(lessonId):
    discussion = discussions.find_one({"lessonId": lessonId})
    return parse_object_id(discussion)


@app.post("/discussions/{lessonId}", response_model=DiscussionOutput, status_code=201)
async def post_discussion(
    lesson: DiscussionInput, authorization: str = Header(None)
) -> Response:
    validation_headers = {
      'Authorization': authorization
    }

    auth_service_url = os.environ.get('AUTH_SERVICE_URL')
    
    # call GO auth service
    auth_service = f"{auth_service_url}/validate"
    auth = requests.get(auth_service, headers=validation_headers)

    if auth.status_code != 200:
      return Response(status_code=401)
    
    auth = auth.json()

    comment = {
        "title": lesson.title,
        "content": lesson.content,
        "createdAt": datetime.datetime.utcnow(),
        "user": auth["username"],
    }
    
    discussions.update_one(
        filter={"lessonId": lesson.lessonId},
        update={"$set": {"lessonId": lesson.lessonId}, "$push": {"comments": comment}},
        upsert=True,
    )
    discussion = discussions.find_one({"lessonId": lesson.lessonId})
    return parse_object_id(discussion)
