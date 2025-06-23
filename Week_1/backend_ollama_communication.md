# Communication Between backend_go and ollama-comps

## 1. How `backend_go` and `ollama-comps` Communicate

### Architecture Overview
- **backend_go** (Go): Main API and business logic server, exposes REST endpoints for the frontend and manages the database.
- **ollama-comps** (Python): Microservice providing LLM/AI functionalities (e.g., text generation, vocabulary generation) via HTTP API.
- **Communication:** The Go backend sends HTTP POST requests to ollama-comps, receives responses, and processes or stores the results in its own database.

### Typical Flow
1. **Trigger:** User action (e.g., requesting new vocabulary) triggers an API call to the Go backend.
2. **Request:** Go backend constructs a request payload and sends it to ollama-comps (e.g., `POST http://ollama-comps:8000/v1/example-service`).
3. **Response:** ollama-comps processes the request (using an LLM or other AI model) and returns a structured response.
4. **Processing:** Go backend parses the response and stores new data (e.g., words, groups) in its database.

---

## 2. Required Changes for Communication

### A. Changes in `backend_go`
- **Add HTTP Client Logic:** Use Go's `net/http` or a library like `resty` to send POST requests to ollama-comps.
- **Define Request/Response Structs:** Mirror the request/response schema expected by ollama-comps (e.g., `ChatCompletionRequest`, `ChatCompletionResponse`).
- **Integrate with Business Logic:** Add service/controller logic to trigger ollama-comps calls when needed (e.g., when importing/generating vocabulary).
- **Error Handling:** Handle errors from ollama-comps gracefully and propagate meaningful messages to the frontend.

### B. Changes in `ollama-comps`
- **API Endpoint Consistency:** Ensure the endpoint (e.g., `/v1/example-service`) is well-documented and stable.
- **Flexible Output:** Support output formats that match the Go backend's database models (e.g., lists of words, groups).
- **Custom Generation Logic:** Optionally, add endpoints or parameters for specific tasks (e.g., "generate vocabulary list for topic X").
- **CORS/Network:** Ensure the service is accessible from the Go backend (correct host/port, Docker networking if needed).

### C. (Optional) Shared Protocol
- Define a shared protocol/schema (e.g., using OpenAPI or JSON Schema) for requests and responses to avoid mismatches.

---

## 3. Functionalities Provided by `ollama-comps`

### A. Core Features
- **Text Generation:** Generate text completions, explanations, or translations using LLMs.
- **Vocabulary Generation:** Given a topic, prompt, or example, generate a list of vocabulary words (with translations, example sentences, etc.).
- **Content Summarization:** Summarize texts or generate study materials.
- **Conversational AI:** Simulate chat or Q&A for language learning.

### B. Database Population (Most Important)
#### How to Use for Database Population
- The Go backend sends a prompt like "Generate 20 beginner Arabic words with English translations and example sentences."
- ollama-comps returns a structured JSON response, e.g.:

```json
{
  "words": [
    {
      "arabic_word": "بيت",
      "english_word": "house",
      "parts": ["noun"],
      "example_sentence": "هذا بيت كبير.",
      "groups": ["Basic Vocabulary"]
    },
    ...
  ]
}
```

- The Go backend parses this response and creates `Word` and `Group` records in its database, matching the models:
  - `Word` model fields: `arabic_word`, `english_word`, `parts`, `groups`, etc.
  - `Group` model fields: `name`, `words`, etc.

#### Mapping Example
- **ollama-comps output:**
```json
{
  "arabic_word": "مدرسة",
  "english_word": "school",
  "parts": ["noun"],
  "groups": ["Education"]
}
```
- **Go Model:**
```go
// Word.go
 type Word struct {
   ArabicWord  string
   EnglishWord string
   Parts       []string
   Groups      []Group
 }
```

### C. Other Potential Features
- **Grammar Explanations:** Generate grammar notes or explanations for given words or sentences.
- **Quiz/Exercise Generation:** Create fill-in-the-blank, multiple-choice, or translation exercises.
- **Pronunciation Guides:** Generate phonetic transcriptions or audio links (if supported).
- **Study Session Content:** Generate content for study sessions or activities.

---

## Summary Table: Model Mapping

| ollama-comps Output Field | Go Model Field         | Notes                        |
|--------------------------|------------------------|------------------------------|
| arabic_word              | Word.ArabicWord        | Required                     |
| english_word             | Word.EnglishWord       | Required                     |
| parts                    | Word.Parts             | Optional (e.g., ["noun"])    |
| groups                   | Word.Groups            | Many-to-many with Group      |
| example_sentence         | (custom, not in model) | Can be added if needed       |

---

## Example Workflow: Populating the Database

1. **Request:**
   - Go backend sends:
     ```json
     {
       "prompt": "Generate 20 beginner Arabic words with English translations and group them by topic."
     }
     ```
2. **Response:**
   - ollama-comps returns:
     ```json
     {
       "words": [
         {"arabic_word": "قلم", "english_word": "pen", "parts": ["noun"], "groups": ["Stationery"]},
         ...
       ]
     }
     ```
3. **Processing:**
   - Go backend iterates over the list, creates `Word` and `Group` records, and links them as per the schema.

---

## Conclusion
- **ollama-comps** can be a powerful backend for generating and importing language learning content, especially vocabulary, directly into the Go backend's database.
- **Integration** requires HTTP client code in Go, a well-defined API contract, and possibly some enhancements to the ollama-comps output format to match the Go models.
- **Extensibility:** The same pattern can be used for other AI-powered features (exercises, explanations, etc.). 