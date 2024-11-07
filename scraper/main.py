from config import SUPABASE_URL, SUPABASE_KEY, SERPAPI_KEY, APP_ID
from scraper.fetch_reviews import fetch_reviews
from classifier.model_utils import ONNXModel
from classifier.classify_reviews import classify_review
from supabase import create_client
from dotenv import load_dotenv

load_dotenv()  # Load environment variables from .env

# Initialize Supabase client
supabase = create_client(SUPABASE_URL, SUPABASE_KEY)

def save_reviews_to_supabase(reviews):
    """Save the list of reviews to Supabase."""
    for review in reviews:
        # Only insert if review_text is present
        if review["review_text"]:
            data = {
                "review_text": review["review_text"],
                "review_date": review["review_date"],
                "rating": review["rating"],
                "category": review.get("category")
            }
            supabase.table("reviews").insert(data).execute()
    print("Reviews saved successfully!")

def main():
    model = ONNXModel("classifier/models/game_review_model.onnx")
    
    # Fetch reviews using SerpAPI
    reviews = fetch_reviews(APP_ID, SERPAPI_KEY, max_reviews=50)
    
    # Classify each review
    for review in reviews:
        category = classify_review(review["review_text"], model)
        review["category"] = category
    
    # Save to Supabase
    save_reviews_to_supabase(reviews)

if __name__ == "__main__":
    main()