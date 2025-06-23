import os
import chromadb
from chromadb.config import Settings

# Set up ChromaDB client (in-memory for now)
client = chromadb.Client(Settings())

# Create (or get) a collection
collection = client.get_or_create_collection("all-my-documents")

# Load all .txt files from local folder
folder_path = "./documents"  # CHANGE THIS if needed
documents = []
ids = []
metadatas = []

for i, filename in enumerate(os.listdir(folder_path)):
    if filename.endswith(".txt"):
        file_path = os.path.join(folder_path, filename)
        with open(file_path, "r", encoding="utf-8") as f:
            content = f.read()
            documents.append(content)
            ids.append(f"doc_{i}")
            metadatas.append({"source": filename})

# Add the documents to the collection
collection.add(
    documents=documents,
    ids=ids,
    metadatas=metadatas
)

print(f"Added {len(documents)} documents from '{folder_path}'.")

# Query example
query = "This is a query document"
results = collection.query(query_texts=[query], n_results=2)

print("\n--- Query Results ---")
for i, result in enumerate(results["documents"][0]):
    print(f"\nResult {i+1}:\n{result}")
