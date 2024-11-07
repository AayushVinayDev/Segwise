from scraper.classifier.classify_reviews import classify_review
from classifier.model_utils import ONNXModel

def test_classify_review():
    model = ONNXModel("game_review_model.onnx")
    category = classify_review("The app keeps crashing.", model)
    assert category in ["Bugs", "Complaints", "Crashes", "Praises", "Other"]
    assert category == "Crashes"  # Adjust based on model's behavior