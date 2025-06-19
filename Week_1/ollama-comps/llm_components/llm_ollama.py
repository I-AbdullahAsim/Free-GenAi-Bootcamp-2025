

import requests
import os

OLLAMA_API_URL = "http://localhost:8008/api/generate"

class OpeaTextGenService:
    def __init__(self, name: str, description: str, config: dict = None):
        self.name = name
        self.description = description
        self.config = config
        self.model_id = os.getenv("LLM_MODEL_ID", "llama3:8b")

    async def invoke(self, input):
        prompt = ""

        if isinstance(input, dict) and "messages" in input:
            messages = input["messages"]
            if isinstance(messages, list):
                prompt = "\n".join([msg["content"] for msg in messages])

        data = {
            "model": self.model_id,
            "prompt": prompt,
            "stream": False
        }

        try:
            res = requests.post(OLLAMA_API_URL, json=data)
            res.raise_for_status()
            output = res.json().get("response", "")
            return {"choices": [{"message": {"content": output}}]}
        except Exception as e:
            return {"error": str(e)}
