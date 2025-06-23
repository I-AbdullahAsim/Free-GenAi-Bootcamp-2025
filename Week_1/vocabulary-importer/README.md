# Vocabulary Importer

A Next.js app for generating and managing English–Arabic vocabulary sets for language learning, powered by LLMs (Groq API, e.g., Llama 3). This tool is designed to help educators and learners quickly create comprehensive, category-based vocabulary lists, and export them for use in games or other educational tools.

---

## Features

- **Generate Vocabulary Sets:**
  - Enter any thematic category (e.g., "Animals", "Kitchen utensils", "Planets").
  - The app uses an LLM to generate a complete, context-aware list of English–Arabic vocabulary items for that category.
  - Each item includes:
    - `english_word`: The English term
    - `arabic_word`: The Arabic translation (with diacritics if useful)
    - `arabic_letters`: Array of individual Arabic letters (auto-generated)
    - `parts`: Array of tags (category, context, singular/plural forms, etc.)

- **Smart Quantity:**
  - The system determines the appropriate number of items based on the category's scope (e.g., 7 for days of the week, 30–50+ for animals).

- **JSON Export:**
  - View and copy the generated vocabulary as JSON, ready for use in typing games or other applications.

- **Send to Seeds:**
  - Save the generated vocabulary and metadata to a backend `seeds` directory for further use (e.g., seeding a database).

- **Modern UI:**
  - Built with React, Next.js, Tailwind CSS, and shadcn/ui for a responsive, user-friendly experience.

---

## Setup & Installation

### 1. **Clone the Repository**

```bash
git clone <your-repo-url>
cd Week_1/vocabulary-importer
```

### 2. **Install Dependencies**

```bash
npm install
# or
yarn install
# or
pnpm install
```

### 3. **Set Up Groq API Key**

This app uses the Groq LLM API via `@ai-sdk/groq`. You must provide your own API key.

- Get your API key from [Groq Console](https://console.groq.com/keys).
- Create a `.env.local` file in the project root:

```
GROQ_API_KEY=your_actual_groq_api_key_here
```

- **Restart the dev server** after adding your key.

### 4. **Run the Development Server**

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
```

- Open [http://localhost:3000](http://localhost:3000) in your browser.

---

## Usage

1. **Enter a Category:**
   - Type a category (e.g., "Animals", "Kitchen utensils") in the input field.
2. **Generate Vocabulary:**
   - Click "Generate Complete Vocabulary Set".
   - The app will call the LLM and display a list of vocabulary items.
3. **View/Copy JSON:**
   - The generated vocabulary appears as a formatted JSON block for easy export.
4. **Send to Seeds:**
   - Click "Send to Seeds" to save the vocabulary and metadata to the backend seeds directory.
   - Files are saved to: `free-genai-bootcamp-2025/Week_1/backend_go/seeds/<category>/`
   - Two files are created:
     - `<category>_vocabulary.json`: The vocabulary data
     - `metadata.json`: Info about the generation
5. As of right now all data inside the seeds folder is generated using this vocabulary_importer
---

## Directory Structure

```
Week_1/vocabulary-importer/
├── app/
│   ├── api/
│   │   ├── generate-vocabulary/route.ts   # LLM-powered vocabulary generation API
│   │   └── send-to-seeds/route.ts         # Save vocabulary to backend seeds
│   ├── page.tsx                           # Main UI page
│   └── ...
├── components/                            # UI components
├── hooks/                                 # React hooks
├── lib/                                   # Utility functions
├── public/                                # Static assets
├── styles/                                # Global styles
├── .env.local                             # (You create this for your API key)
├── package.json                           # Project dependencies
└── ...
```

---

## Troubleshooting

- **No vocabulary generated?**
  - Make sure your `.env.local` file contains a valid `GROQ_API_KEY`.
  - Restart the dev server after adding or changing the key.

- **Cannot write to seeds directory?**
  - Ensure the backend seeds path exists and is writable by your server process.
  - If running in WSL2, check permissions and that you are looking in the correct file system.
  - This feature will not work on Vercel or other serverless/cloud deployments (local/server only).

- **Still having issues?**
  - Check the terminal for error logs when you click "Send to Seeds".
  - Add debug logs to the API handler if needed.

---

