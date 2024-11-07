import os
from dotenv import load_dotenv

load_dotenv()  # Load environment variables from a .env file

SUPABASE_URL = os.getenv("SUPABASE_URL")
SUPABASE_KEY = os.getenv("SUPABASE_KEY")
APP_ID = os.getenv("APP_ID")  # Google Play app ID for scraping reviews
SERPAPI_KEY = os.getenv("SERPAPI_KEY")  # Google Play app ID for scraping reviews
