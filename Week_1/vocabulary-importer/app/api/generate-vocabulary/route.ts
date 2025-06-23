import { type NextRequest, NextResponse } from "next/server"
import { generateText } from "ai"
import { groq } from "@ai-sdk/groq"
import { z } from "zod"

// ---------- Zod schema ----------
const VocabularySchema = z.array(
  z.object({
    english_word: z.string(),
    arabic_word: z.string(),
    arabic_letters: z.array(z.string()),
    parts: z.array(z.string()),
  }),
)

type Vocabulary = z.infer<typeof VocabularySchema>

// ---------- Helper to break Arabic text into individual letters ----------
function breakIntoArabicLetters(arabicText: string): string[] {
  // Remove diacritics and spaces, then split into individual characters
  const cleanText = arabicText.replace(/[\u064B-\u065F\u0670\u06D6-\u06ED\s]/g, "")
  return Array.from(cleanText).filter((char) => char.trim() !== "")
}

// ---------- Helper to clean / repair model output ----------
function cleanJson(raw: string): string {
  return (
    raw
      // strip markdown fences / language tags
      .replace(/```json|```/gi, "")
      // trim whitespace
      .trim()
      // remove trailing commas before ] or }
      .replace(/,\s*([\]}])/g, "$1")
  )
}

// ---------- Core generation logic with retries ----------
async function generateVocabulary(category: string, attempts = 3): Promise<Vocabulary> {
  for (let i = 1; i <= attempts; i++) {
    const { text } = await generateText({
      model: groq("llama3-70b-8192"),
      temperature: 0.2,
      system:
        "You are an API that must return ONLY valid JSON (no markdown). " +
        "The JSON must be an array where each element matches the provided schema.",
      prompt: `
Generate English–Arabic vocabulary items for the thematic category «${category}».

IMPORTANT: Generate ALL the common, essential, and frequently used items that belong to this category. 

Guidelines for quantity:
- If it's a broad category (like "Animals", "Food", "Transportation"), include 30-50+ items
- If it's a specific category (like "Days of the week", "Months", "Primary colors"), include all of them (7, 12, 6 respectively)
- If it's a medium category (like "Kitchen utensils", "School supplies"), include 15-25 items
- If it's a very specific category (like "Planets in solar system"), include all 8-9 items
- Always prioritize completeness and usefulness over arbitrary numbers

For each item include:
  • english_word  – English term
  • arabic_word   – Accurate Arabic translation (with diacritics if useful)
  • arabic_letters – Array of individual Arabic letters (will be auto-generated, so just include empty array [])
  • parts         – Array (3–6 strings) that includes:
        – Category/type (e.g. "fruit", "animal", "utensil")
        – Context (e.g. "food", "transportation", "education")
        – A "singular: ..." Arabic form
        – A "plural: ..." Arabic form

Example format:
[
  {
    "english_word": "apple",
    "arabic_word": "تفاحة",
    "arabic_letters": [],
    "parts": ["fruit", "food", "singular: تفاحة", "plural: تفاحات"]
  }
]

Think about what items are truly essential and commonly used in the category "${category}". 
Include ALL the important ones - don't limit yourself to exactly 20 items.

Return ONLY the JSON array with all relevant items for this category.
      `,
    })

    try {
      const parsed = JSON.parse(cleanJson(text))

      // Add arabic_letters breakdown for each item
      const withLetters = parsed.map((item: any) => ({
        ...item,
        arabic_letters: breakIntoArabicLetters(item.arabic_word),
      }))

      const validated = VocabularySchema.parse(withLetters)
      return validated
    } catch (err) {
      console.warn(`Attempt ${i} failed JSON validation`, err)
      if (i === attempts) throw err
    }
  }
  // Should never reach here
  throw new Error("Exhausted all attempts to generate valid vocabulary JSON")
}

// ---------- Route handler ----------
export async function POST(request: NextRequest) {
  try {
    const { category } = await request.json()

    if (!category || typeof category !== "string") {
      return NextResponse.json({ error: "Category is required and must be a string" }, { status: 400 })
    }

    const vocabulary = await generateVocabulary(category)

    return NextResponse.json({ vocabulary })
  } catch (error) {
    console.error("Error generating vocabulary:", error)
    return NextResponse.json({ error: "Failed to generate vocabulary" }, { status: 500 })
  }
}
