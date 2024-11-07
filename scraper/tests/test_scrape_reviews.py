from scraper.scraper.scrape_reviews import scrape_reviews
from config import APP_ID
from datetime import datetime, timedelta

def test_scrape_reviews():
    reviews = scrape_reviews(APP_ID, max_reviews_per_day=15, days=7)
    assert len(reviews) > 0
    for review in reviews:
        assert "review_text" in review
        assert "review_date" in review
        assert datetime.strptime(review["review_date"], "%Y-%m-%d") >= datetime.now() - timedelta(days=7)