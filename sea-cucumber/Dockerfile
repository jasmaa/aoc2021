FROM python:3.6-alpine
WORKDIR /app

RUN pip install -U pip
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .

ENTRYPOINT ["./entrypoint.sh"]