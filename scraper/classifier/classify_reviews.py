CATEGORY_PRIORITY = ["Bugs", "Complaints", "Crashes", "Praises", "Other"]

def classify_review(text, model):
    # Skip classification if text is empty or None
    if not text:
        return "Other"  # Default category for empty reviews

    # Proceed with classification
    category_index = model.predict(text)
    return CATEGORY_PRIORITY[category_index]