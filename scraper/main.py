from config import SUPABASE_URL, SUPABASE_KEY, APP_ID
from scraper.fetch_reviews import fetch_reviews
from classifier.model_utils import ONNXModel
from classifier.classify_reviews import classify_review
from supabase import create_client
from dotenv import load_dotenv

load_dotenv()  # Load environment variables from .env

# Initialize Supabase client
supabase = create_client(SUPABASE_URL, SUPABASE_KEY)

def save_reviews_to_supabase(reviews):
    """Save the list of reviews to Supabase, avoiding duplicates."""
    if not reviews:
        print("No new reviews to save.")
        return

    # Fetch existing reviews' text and dates
    try:
        response = supabase.table("reviews").select("review_text", "review_date").execute()
        existing_reviews = set(
            (review["review_text"], review["review_date"]) for review in response.data
        )
    except Exception as e:
        print(f"Error fetching existing reviews: {e}")
        return

    new_reviews = []
    for review in reviews:
        if not review["review_text"]:
            continue

        key = (review["review_text"], review["review_date"])
        if key in existing_reviews:
            print("Review already exists. Skipping.")
            continue
        else:
            new_reviews.append({
                "review_text": review["review_text"],
                "review_date": review["review_date"],
                "rating": review["rating"],
                "category": review.get("category")
            })

    if not new_reviews:
        print("No new reviews to save after checking for duplicates.")
        return

    # Insert new reviews into Supabase
    try:
        response = supabase.table("reviews").insert(new_reviews).execute()
        print(f"Inserted {len(response.data)} new reviews successfully.")
    except Exception as e:
        print(f"Error inserting reviews: {e}")

def main():
    model = ONNXModel("classifier/models/game_review_model.onnx")

    # Fetch new reviews from the last 7 days
    print("Fetching new reviews from the last 7 days...")
    reviews = fetch_reviews(APP_ID, days=7, max_reviews=100)
    print(f"Total new reviews fetched: {len(reviews)}")

    if not reviews:
        print("No new reviews to process.")
        return

    # Classify reviews
    print("Classifying reviews...")
    for idx, review in enumerate(reviews, 1):
        try:
            # Ensure 'review_text' key exists
            review_text = review["review_text"]
            if not review_text:
                raise ValueError("Empty review text.")

            category = classify_review(review_text, model)
            review["category"] = category
        except KeyError as e:
            print(f"Error in review {idx}: Missing key {e}. Skipping review.")
            review["category"] = "Unknown"
        except ValueError as e:
            print(f"Value error in review {idx}: {e}. Setting category to 'Unknown'.")
            review["category"] = "Unknown"
        except RuntimeError as e:
            print(f"Runtime error in review {idx}: {e}. Setting category to 'Unknown'.")
            review["category"] = "Unknown"
        except Exception as e:
            # Optional: Catch other exceptions and log them
            print(f"Unexpected error in review {idx}: {e}. Setting category to 'Unknown'.")
            review["category"] = "Unknown"

        if idx % 10 == 0 or idx == len(reviews):
            print(f"Classified {idx}/{len(reviews)} reviews")

    # Save to Supabase
    print("Saving reviews to Supabase...")
    save_reviews_to_supabase(reviews)

if __name__ == "__main__":
    main()
