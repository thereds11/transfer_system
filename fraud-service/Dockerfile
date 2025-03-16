FROM python:3.11-slim
WORKDIR /app
COPY . .
RUN pip install --no-cache-dir -r requirements.txt watchdog
# Run with watchmedo
CMD ["watchmedo", "auto-restart", "--patterns=*.py", "--recursive", "--", "python", "/app/main.py"]