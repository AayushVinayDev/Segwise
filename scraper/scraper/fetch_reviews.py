import requests

def fetch_reviews(app_id, serpapi_key, max_reviews=50):
    """Fetches reviews from Google Play Store using SerpApi."""
    url = "https://serpapi.com/search.json"
    params = {
        "engine": "google_play_product",
        "product_id": app_id,
        "store": "apps",  # Add store parameter
        "all_reviews": "true",
        "num": max_reviews,
        "api_key": serpapi_key
    }
    
    response = requests.get(url, params=params, timeout=30)
    if response.ok:
        data = response.json()
        reviews = data.get("reviews", [])
        return [
            {
                "review_text": review.get("snippet", "").strip() or None,  # Use "snippet" field for review text
                "review_date": review.get("date"),
                "rating": int(review.get("rating", 0)),  # Ensure rating is an integer
            }
            for review in reviews if review.get("snippet")  # Filter out empty reviews
        ]
    else:
        print("Error fetching reviews:", response.status_code, response.text)
        response.raise_for_status()