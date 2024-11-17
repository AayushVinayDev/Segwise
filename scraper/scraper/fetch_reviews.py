from google_play_scraper import reviews, Sort
from datetime import datetime, timedelta
import time
import requests

def fetch_reviews(app_id, days=7, max_reviews=100):
    """
    Fetches up to max_reviews reviews from the last 'days' days from Google Play Store using google-play-scraper.
    """
    all_reviews = []
    today = datetime.now()
    cutoff_date = today - timedelta(days=days)
    continuation_token = None

    while True:
        try:
            # Fetch reviews using google-play-scraper
            result, continuation_token = reviews(
                app_id,
                lang='en',
                country='us',
                sort=Sort.NEWEST,
                count=200,  # Max number per batch
                continuation_token=continuation_token
            )
        except requests.exceptions.RequestException as e:
            print(f"Network error occurred: {e}")
            break

        if not result:
            break  # No more reviews

        for review in result:
            review_date = review['at']  # This is a datetime object
            if review_date < cutoff_date:
                # Review is older than the cutoff date
                return all_reviews[:max_reviews]

            review_data = {
                "author_name": review["userName"],
                "review_text": review["content"],
                "review_date": review_date.strftime("%Y-%m-%d"),
                "rating": review["score"],
            }
            all_reviews.append(review_data)

            if len(all_reviews) >= max_reviews:
                return all_reviews[:max_reviews]

        if not continuation_token:
            break  # No more pages to fetch

        # Add a delay to avoid potential rate limits
        time.sleep(1)  # Sleep for 1 second between requests

    return all_reviews[:max_reviews]
