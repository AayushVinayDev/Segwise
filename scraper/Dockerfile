FROM python:3.9-slim

WORKDIR /app

# Copy files
COPY . .

# Install dependencies
RUN pip install -r requirements.txt

# Run the scraper
CMD ["python", "main.py"]