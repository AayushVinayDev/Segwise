# Review Classification Project

This project is a review classification system designed to scrape, classify, store, and retrieve reviews for a specific app from the Google Play Store. The project is divided into two main parts: a Python-based scraper and classifier, and a Golang-based API with an HTMX frontend for querying and displaying review data.

## Project Structure

project-root/
├── backend-api/
│   ├── cmd/
│   ├── config/
│   ├── controllers/
│   ├── db/
│   ├── models/
│   ├── routes/
│   ├── services/
│   ├── static/
│   ├── tests/
│   ├── .env
│   ├── .env.example
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
└── scraper/
    ├── classifier/
    ├── scraper/
    ├── tests/
    ├── __init__.py
    ├── .env
    ├── .env.example
    ├── config.py
    ├── Dockerfile
    ├── main.py
    └── requirements.txt

## Model Accuracy and Improvements

The review classification model uses `distilbert-base-uncased`, which is optimized for low production costs. Our model was trained with a custom dataset of 1000 reviews and achieved high accuracy during training. However, accuracy could be improved with additional labeled data and extended fine-tuning, especially for capturing nuances in user reviews.

- **Model**: DistilBERT (distilbert-base-uncased) for fast, lightweight inference.
- **Model Accuracy**: Achieved good accuracy with the 1000-row custom dataset. Additional data could further improve performance.
- **ONNX Model Size**: 256MB.

You can access the Colab notebook used for model training and ONNX conversion here: [Colab Notebook Link](https://colab.research.google.com/drive/1Aj5pLS6y797JqGMblBjKO0WC3AyMitma?usp=sharing).

## Project Workflow

### Data Scraping and Classification (Python)
The Python scraper uses SerpAPI to retrieve reviews from the Google Play Store. Each review is classified into one of five categories using our ONNX model and then stored in a Supabase-hosted PostgreSQL database.

- **Frequency**: Runs every 24 hours.
- **Limitations**: Free tier restrictions on SerpAPI limit the number of reviews that can be scraped daily.

### Backend API (Golang) and HTMX Frontend
The Golang backend API retrieves categorized review data from Supabase, performs necessary operations, and serves it to the HTMX frontend for visualization. The HTMX UI allows users to view reviews by category and date.

- **Endpoints**:
    - **`/reviews`**: Fetches reviews by category and date.
    - **`/trend`**: Shows a 7-day trend of review counts per category.

## API Sample Requests and Responses

### 1. `/reviews` Endpoint
**Request**:
```http
GET /reviews?category=Praises&date=2024-11-01

[
    {
        "id": "b102759c-f4db-4fc6-9b7b-57757870f440",
        "review_text": "I like this game but I don't love it...",
        "review_date": "2024-11-01",
        "rating": 3,
        "category": "Praises"
    },
    {
        "id": "7077f260-77aa-48c3-96cc-e44a5d8a916f",
        "review_text": "I love this game and I play it all the time...",
        "review_date": "2024-11-01",
        "rating": 3,
        "category": "Praises"
    }
]

GET /trend?category=Praises&date=2024-11-01

[
    {"date": "2024-10-25", "count": 10},
    {"date": "2024-10-26", "count": 15},
    ...
]
```

## Local Setup Instructions

### Backend API:

1. **Prerequisites**: Install [Docker](https://www.docker.com/) and [Go](https://golang.org/).
2. **Clone the repository** and navigate to the `backend-api` directory.
3. **Set up a `.env` file** with the necessary environment variables:

    ```makefile
    SUPABASE_URL=<Your_Supabase_URL>
    SUPABASE_KEY=<Your_Supabase_API_Key>
    SERPAPI_KEY=<Your_SerpAPI_Key>
    ```

4. **Build and Run**:

    ```bash
    docker build -t backend-api .
    docker run -p 8080:8080 --env-file .env backend-api
    ```

### Scraper:

1. **Prerequisites**: Install Python 3 and dependencies from `requirements.txt`.
2. **Set up a virtual environment**:

    ```bash
    python3 -m venv venv
    source venv/bin/activate  # On Windows, use venv\Scripts\activate
    ```

3. **Install dependencies**:

    ```bash
    pip install -r requirements.txt
    ```

4. **Set up environment variables**: Copy the provided `env.example` file to a new `.env` file and fill in your specific keys.

5. **Run the scraper**:

    ```bash
    python main.py
    ```

This will scrape reviews, classify them, and store the data in your Supabase database.

### Access the UI:

- Visit [http://localhost:8080](http://localhost:8080) to interact with the HTMX UI.

---

### Deployed Solution

You can access the deployed solution here: [Deployed Link](https://segwise-cdcx.onrender.com/)

---

## Production Cost Estimates (Example for AWS)

| Service                 | Cost Estimate                       |
|-------------------------|-------------------------------------|
| **ONNX Model Hosting**  | $5/month (256MB)                   |
| **PostgreSQL RDS**      | $15/month (Free tier if under 1GB) |
| **Golang Backend (24/7)** | $10/month                        |
| **Static File Hosting** | $0 (served by the backend)         |
| **SerpAPI (Data Scraping)** | Free tier or $50/month for paid tier |

**Total Estimated Monthly Cost**: ~ $30 on the free tier, with SerpAPI and PostgreSQL usage on a paid tier bringing it to ~$80.

---

## Future Improvements

- **Enhanced Model Accuracy**: Using a larger dataset and additional training would improve classification performance.
- **Expanded API Functionality**: Adding filtering and pagination for more comprehensive review browsing.

