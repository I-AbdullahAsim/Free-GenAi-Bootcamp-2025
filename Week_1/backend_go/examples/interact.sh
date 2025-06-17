#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}Starting API interaction examples...${NC}\n"

# Function to check response
check_response() {
    if [[ $1 == *"error"* ]] || [[ $1 == *"false"* ]]; then
        echo -e "${RED}Error: $1${NC}"
        return 1
    else
        echo -e "${GREEN}Success: $1${NC}"
        return 0
    fi
}

# 1. Create a word group
echo -e "${GREEN}1. Creating a word group...${NC}"
GROUP_RESPONSE=$(curl -s -X POST http://localhost:8080/api/groups \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Basic Arabic Words"
  }')
GROUP_ID=$(echo $GROUP_RESPONSE | grep -o '"ID":[0-9]*' | cut -d':' -f2)
check_response "$GROUP_RESPONSE"
echo

# 2. Add some words
echo -e "${GREEN}2. Adding words...${NC}"
WORD1_RESPONSE=$(curl -s -X POST http://localhost:8080/api/words \
  -H "Content-Type: application/json" \
  -d '{
    "arabic_word": "مرحبا",
    "english_word": "Hello",
    "parts": ["greeting", "singular: مرحبا", "plural: مرحبا"]
  }')
WORD1_ID=$(echo $WORD1_RESPONSE | grep -o '"ID":[0-9]*' | cut -d':' -f2)
check_response "$WORD1_RESPONSE"
echo

WORD2_RESPONSE=$(curl -s -X POST http://localhost:8080/api/words \
  -H "Content-Type: application/json" \
  -d '{
    "arabic_word": "شكرا",
    "english_word": "Thank you",
    "parts": ["gratitude", "singular: شكرا", "plural: شكرا"]
  }')
WORD2_ID=$(echo $WORD2_RESPONSE | grep -o '"ID":[0-9]*' | cut -d':' -f2)
check_response "$WORD2_RESPONSE"
echo

# 3. Create a study activity
echo -e "${GREEN}3. Creating a study activity...${NC}"
ACTIVITY_RESPONSE=$(curl -s -X POST http://localhost:8080/api/study-activities \
  -H "Content-Type: application/json" \
  -d "{
    \"group_id\": $GROUP_ID,
    \"name\": \"Basic Review\"
  }")
ACTIVITY_ID=$(echo $ACTIVITY_RESPONSE | grep -o '"ID":[0-9]*' | cut -d':' -f2)
check_response "$ACTIVITY_RESPONSE"
echo

# 4. Start a study session
echo -e "${GREEN}4. Starting a study session...${NC}"
SESSION_RESPONSE=$(curl -s -X POST http://localhost:8080/api/study-sessions \
  -H "Content-Type: application/json" \
  -d "{
    \"group_id\": $GROUP_ID,
    \"activity_id\": $ACTIVITY_ID
  }")
SESSION_ID=$(echo $SESSION_RESPONSE | grep -o '"ID":[0-9]*' | cut -d':' -f2)
check_response "$SESSION_RESPONSE"
echo

# 5. Record some word reviews
echo -e "${GREEN}5. Recording word reviews...${NC}"
REVIEW1_RESPONSE=$(curl -s -X POST http://localhost:8080/api/study-sessions/$SESSION_ID/words/$WORD1_ID/review \
  -H "Content-Type: application/json" \
  -d '{
    "correct": true
  }')
check_response "$REVIEW1_RESPONSE"
echo

REVIEW2_RESPONSE=$(curl -s -X POST http://localhost:8080/api/study-sessions/$SESSION_ID/words/$WORD2_ID/review \
  -H "Content-Type: application/json" \
  -d '{
    "correct": false
  }')
check_response "$REVIEW2_RESPONSE"
echo

# 6. Check study progress
echo -e "${GREEN}6. Checking study progress...${NC}"
PROGRESS_RESPONSE=$(curl -s http://localhost:8080/api/dashboard/study_progress)
check_response "$PROGRESS_RESPONSE"
echo

# 7. View quick stats
echo -e "${GREEN}7. Viewing quick stats...${NC}"
STATS_RESPONSE=$(curl -s http://localhost:8080/api/dashboard/quick_stats)
check_response "$STATS_RESPONSE"
echo

echo -e "\n${BLUE}API interaction examples completed!${NC}" 