"use client"

import type React from "react"

import { useState } from "react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Label } from "@/components/ui/label"
import { Loader2, FolderPlus, CheckCircle, Keyboard, Info } from "lucide-react"
import { useToast } from "@/hooks/use-toast"

interface VocabularyItem {
  english_word: string
  arabic_word: string
  arabic_letters: string[]
  parts: string[]
}

export default function VocabularyImporter() {
  const [category, setCategory] = useState("")
  const [vocabulary, setVocabulary] = useState<VocabularyItem[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [isSending, setIsSending] = useState(false)
  const [jsonOutput, setJsonOutput] = useState("")
  const { toast } = useToast()

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!category.trim()) return

    setIsLoading(true)
    try {
      const response = await fetch("/api/generate-vocabulary", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ category: category.trim() }),
      })

      if (!response.ok) {
        throw new Error("Failed to generate vocabulary")
      }

      const data = await response.json()
      setVocabulary(data.vocabulary)
      setJsonOutput(JSON.stringify(data.vocabulary, null, 2))

      toast({
        title: "Success!",
        description: `Generated ${data.vocabulary.length} vocabulary items for "${category}"`,
      })
    } catch (error) {
      console.error("Error generating vocabulary:", error)
      toast({
        title: "Error",
        description: "Failed to generate vocabulary. Please try again.",
        variant: "destructive",
      })
    } finally {
      setIsLoading(false)
    }
  }

  const sendToSeeds = async () => {
    if (!jsonOutput || !category) return

    setIsSending(true)
    try {
      const response = await fetch("/api/send-to-seeds", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          category: category.trim(),
          vocabulary: vocabulary,
        }),
      })

      if (!response.ok) {
        throw new Error("Failed to send to seeds")
      }

      const data = await response.json()

      toast({
        title: "Success!",
        description: `Vocabulary saved to: ${data.path}`,
      })
    } catch (error) {
      console.error("Error sending to seeds:", error)
      toast({
        title: "Error",
        description: "Failed to send to seeds directory.",
        variant: "destructive",
      })
    } finally {
      setIsSending(false)
    }
  }

  const exampleCategories = [
    { name: "Days of the week", expected: "7 items" },
    { name: "Months of the year", expected: "12 items" },
    { name: "Primary colors", expected: "6-8 items" },
    { name: "Kitchen utensils", expected: "15-25 items" },
    { name: "Animals", expected: "30-50+ items" },
    { name: "Planets in solar system", expected: "8-9 items" },
    { name: "Body parts", expected: "25-40 items" },
    { name: "School subjects", expected: "10-15 items" },
  ]

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 p-4">
      <div className="max-w-6xl mx-auto space-y-6">
        <div className="text-center py-8">
          <h1 className="text-4xl font-bold text-gray-900 mb-2">Smart Vocabulary Language Importer</h1>
          <p className="text-lg text-gray-600">
            Generate complete English-Arabic vocabulary sets based on category scope
          </p>
        </div>

        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Keyboard className="h-5 w-5" />
              Generate Complete Vocabulary Set
            </CardTitle>
            <CardDescription>
              Enter any thematic category - the system will automatically determine the appropriate number of items
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="bg-blue-50 p-4 rounded-lg border border-blue-200">
              <div className="flex items-start gap-2">
                <Info className="h-5 w-5 text-blue-600 mt-0.5 flex-shrink-0" />
                <div>
                  <h4 className="font-medium text-blue-900 mb-2">How it works:</h4>
                  <p className="text-sm text-blue-800 mb-3">
                    The system intelligently determines how many items to generate based on the category's natural
                    scope:
                  </p>
                  <div className="grid grid-cols-1 md:grid-cols-2 gap-2 text-xs">
                    {exampleCategories.map((example, index) => (
                      <div key={index} className="flex justify-between bg-white/50 px-2 py-1 rounded">
                        <span className="font-medium">"{example.name}"</span>
                        <span className="text-blue-700">→ {example.expected}</span>
                      </div>
                    ))}
                  </div>
                </div>
              </div>
            </div>

            <form onSubmit={handleSubmit} className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="category">Thematic Category</Label>
                <Input
                  id="category"
                  type="text"
                  placeholder="e.g., Bedroom, Kitchen utensils, Animals, Planets, Colors..."
                  value={category}
                  onChange={(e) => setCategory(e.target.value)}
                  disabled={isLoading}
                />
              </div>
              <Button type="submit" disabled={isLoading || !category.trim()} className="w-full">
                {isLoading ? (
                  <>
                    <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    Generating Complete Vocabulary Set...
                  </>
                ) : (
                  "Generate Complete Vocabulary Set"
                )}
              </Button>
            </form>
          </CardContent>
        </Card>

        {vocabulary.length > 0 && (
          <div className="grid lg:grid-cols-2 gap-6">
            <Card>
              <CardHeader>
                <CardTitle className="flex items-center gap-2">
                  <CheckCircle className="h-5 w-5 text-green-600" />
                  Generated Vocabulary ({vocabulary.length} items)
                </CardTitle>
                <CardDescription>
                  Complete vocabulary set for "{category}" with Arabic letter breakdowns
                </CardDescription>
              </CardHeader>
              <CardContent>
                <div className="space-y-4 max-h-96 overflow-y-auto">
                  {vocabulary.map((item, index) => (
                    <div key={index} className="p-4 bg-gray-50 rounded-lg border">
                      <div className="flex justify-between items-start mb-3">
                        <span className="font-semibold text-gray-900 text-lg">{item.english_word}</span>
                        <span className="text-xl font-medium text-blue-600">{item.arabic_word}</span>
                      </div>

                      {/* Arabic Letters Breakdown */}
                      <div className="mb-3">
                        <span className="text-sm font-medium text-gray-700 block mb-1">Arabic Letters:</span>
                        <div className="flex flex-wrap gap-1">
                          {item.arabic_letters.map((letter, letterIndex) => (
                            <span
                              key={letterIndex}
                              className="px-2 py-1 bg-green-100 text-green-800 text-sm rounded border font-mono"
                            >
                              {letter}
                            </span>
                          ))}
                        </div>
                      </div>

                      {/* Parts/Tags */}
                      <div className="flex flex-wrap gap-1">
                        {item.parts.map((part, partIndex) => (
                          <span key={partIndex} className="px-2 py-1 bg-blue-100 text-blue-800 text-xs rounded-full">
                            {part}
                          </span>
                        ))}
                      </div>
                    </div>
                  ))}
                </div>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle>JSON Output for Typing Game</CardTitle>
                <CardDescription>
                  Complete dataset with {vocabulary.length} items ready for your typing game
                </CardDescription>
              </CardHeader>
              <CardContent className="space-y-4">
                <Textarea
                  value={jsonOutput}
                  readOnly
                  className="min-h-96 font-mono text-sm"
                  placeholder="Generated JSON will appear here..."
                />
                {jsonOutput && (
                  <Button onClick={sendToSeeds} className="w-full" variant="outline" disabled={isSending}>
                    {isSending ? (
                      <>
                        <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                        Sending to Seeds...
                      </>
                    ) : (
                      <>
                        <FolderPlus className="mr-2 h-4 w-4" />
                        Send to Seeds
                      </>
                    )}
                  </Button>
                )}
              </CardContent>
            </Card>
          </div>
        )}
      </div>
    </div>
  )
}
