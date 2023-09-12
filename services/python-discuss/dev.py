import uvicorn
import os

host = os.environ.get('HOST', '0.0.0.0')
port = os.environ.get('PORT', 5000)

def start():
    uvicorn.run("src.main:app", host=host, port=int(port), reload=True)