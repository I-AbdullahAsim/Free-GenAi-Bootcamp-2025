from comps.cores.proto.api_protocol import (
    ChatCompletionRequest,
    ChatCompletionResponse,
    ChatCompletionResponseChoice,
    ChatMessage,
    UsageInfo
)

from starlette.responses import StreamingResponse
import json

from fastapi import HTTPException
from comps.cores.mega.constants import ServiceType, ServiceRoleType
from comps import MicroService, ServiceOrchestrator

import os

# Environment variables with defaults
EMBEDDING_SERVICE_HOST_IP = os.getenv("EMBEDDING_SERVICE_HOST_IP", "0.0.0.0")
EMBEDDING_SERVICE_PORT = os.getenv("EMBEDDING_SERVICE_HOST_PORT", 6000)
LLM_SERVICE_HOST_IP = os.getenv("LLM_SERVICE_HOST_IP", "0.0.0.0")
LLM_SERVICE_PORT = os.getenv("LLM_SERVICE_HOST_PORT", 9000)

class ExampleService:
    def __init__(self, host="0.0.0.0", port=8000):
        print("Hello")
        self.host = host
        self.port = port
        self.endpoint = "/v1/example-service"
        self.megaservice = ServiceOrchestrator()

    def add_remote_service(self):
        llm = MicroService(
            name="llm",
            host=LLM_SERVICE_HOST_IP,
            port=LLM_SERVICE_PORT,
            endpoint="/api/generate",
            use_remote_service=True,
            service_type=ServiceType.LLM,
        )
        self.megaservice.add(llm)

    def start(self):
        self.service = MicroService(
            self.__class__.__name__,
            service_role=ServiceRoleType.MEGASERVICE,
            host=self.host,
            port=self.port,
            endpoint=self.endpoint,
            input_datatype=ChatCompletionRequest,
            output_datatype=ChatCompletionResponse,
        )
        self.service.add_route(self.endpoint, self.handle_request, methods=["POST"])
        self.service.start()

    async def handle_request(self, request: ChatCompletionRequest) -> ChatCompletionResponse:
        try:
            # Extract the latest user message as prompt
            messages = request.messages or []
            user_messages = [m.content for m in messages if m.role == "user"]
            prompt = user_messages[-1] if user_messages else ""

            payload = {
                "model": request.model or "llama3.2:1b",
                "prompt": prompt,
                "stream": False
            }

            result = await self.megaservice.schedule(payload)

            if isinstance(result, tuple):
                result_dict = result[0]
            else:
                result_dict = result

            response_stream = list(result_dict.values())[0]

            if isinstance(response_stream, StreamingResponse):
                content = b""
                async for chunk in response_stream.body_iterator:
                    content += chunk

                decoded = content.decode().strip()

                if not decoded:
                    raise HTTPException(status_code=500, detail="LLM service returned empty response.")

                try:
                    print("LLM raw response:", decoded)
                    llm_response_json = json.loads(decoded)
                except json.JSONDecodeError as e:
                    raise HTTPException(status_code=500, detail=f"Invalid JSON from LLM service: {e}")

                generated_text = llm_response_json.get("response", "[No 'response' key in LLM output]")
            else:
                generated_text = str(response_stream)

            return ChatCompletionResponse(
                model=request.model or "example-model",
                choices=[
                    ChatCompletionResponseChoice(
                        index=0,
                        message=ChatMessage(role="assistant", content=generated_text),
                        finish_reason="stop"
                    )
                ],
                usage=UsageInfo(
                    promt_token=0,
                    completion_tokens=0,
                    total_tokens=0
                )
            )
        except Exception as e:
            raise HTTPException(status_code=500, detail=f"500: {str(e)}")


# Run the service
example = ExampleService()
example.add_remote_service()
example.start()
