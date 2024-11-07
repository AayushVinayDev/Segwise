# classifier/model_utils.py
import onnxruntime
from transformers import DistilBertTokenizer

class ONNXModel:
    def __init__(self, model_path="classifier/models/game_review_model.onnx"):
        self.session = onnxruntime.InferenceSession(model_path)
        self.tokenizer = DistilBertTokenizer.from_pretrained("distilbert-base-uncased")

    def predict(self, text):
        # Tokenize the input text with max_length=5 to match model's expected input dimensions
        tokens = self.tokenizer(text, return_tensors="np", padding="max_length", truncation=True, max_length=5)
        
        # Convert inputs to int64 to match model's requirements
        inputs = {k: v.astype('int64') for k, v in tokens.items()}
        
        # Run the model and get logits
        logits = self.session.run(None, inputs)[0]
        return logits.argmax(axis=-1).item()  # Returns the category index
