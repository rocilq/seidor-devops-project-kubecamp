FROM mcr.microsoft.com/devcontainers/python:3.10-bullseye

RUN pip install poetry

WORKDIR /app

COPY . .
RUN poetry install

CMD ["poetry", "run", "dev"]