
import { useParams } from "react-router-dom";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { ArrowLeft, Volume2, BookOpen, Users } from "lucide-react";
import { Link } from "react-router-dom";

const WordDetail = () => {
  const { id } = useParams();

  // Mock data - in a real app, this would be fetched based on the ID
  const word = {
    id: parseInt(id || "1"),
    arabic: "بيت",
    english: "house",
    pronunciation: "bayt",
    category: "Family & Home",
    difficulty: "Beginner",
    learned: true,
    definition: "A building for human habitation, especially one that is lived in by a family or small group of people.",
    examples: [
      { arabic: "هذا بيتي", english: "This is my house", pronunciation: "hadha bayti" },
      { arabic: "البيت كبير", english: "The house is big", pronunciation: "al-bayt kabir" },
      { arabic: "في البيت", english: "In the house", pronunciation: "fi al-bayt" },
    ],
    relatedWords: [
      { arabic: "غرفة", english: "room", pronunciation: "ghurfa" },
      { arabic: "باب", english: "door", pronunciation: "bab" },
      { arabic: "نافذة", english: "window", pronunciation: "nafidha" },
    ],
    groups: ["Basic Vocabulary", "Family & Home"],
  };

  const getDifficultyColor = (difficulty: string) => {
    switch (difficulty) {
      case "Beginner": return "bg-green-100 text-green-700";
      case "Intermediate": return "bg-yellow-100 text-yellow-700";
      case "Advanced": return "bg-red-100 text-red-700";
      default: return "bg-gray-100 text-gray-700";
    }
  };

  return (
    <div className="space-y-8">
      <div className="flex items-center space-x-4">
        <Link to="/words">
          <Button variant="outline" size="sm">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Words
          </Button>
        </Link>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <div className="lg:col-span-2 space-y-6">
          {/* Main Word Info */}
          <Card>
            <CardHeader>
              <div className="flex items-start justify-between">
                <div className="space-y-2">
                  <CardTitle className="text-4xl font-bold text-blue-600">
                    {word.arabic}
                  </CardTitle>
                  <CardDescription className="text-2xl font-medium">
                    {word.english}
                  </CardDescription>
                  <p className="text-lg text-gray-500 italic">/{word.pronunciation}/</p>
                </div>
                <div className="flex items-center space-x-2">
                  {word.learned && (
                    <div className="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                      <span className="text-green-600 font-bold">✓</span>
                    </div>
                  )}
                  <Button variant="outline" size="sm">
                    <Volume2 className="h-4 w-4" />
                  </Button>
                </div>
              </div>
            </CardHeader>
            <CardContent>
              <div className="flex items-center space-x-4 mb-4">
                <Badge variant="outline">{word.category}</Badge>
                <Badge className={getDifficultyColor(word.difficulty)}>
                  {word.difficulty}
                </Badge>
              </div>
              <p className="text-gray-700">{word.definition}</p>
            </CardContent>
          </Card>

          {/* Examples */}
          <Card>
            <CardHeader>
              <CardTitle>Example Sentences</CardTitle>
              <CardDescription>See how this word is used in context</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {word.examples.map((example, index) => (
                  <div key={index} className="border-l-4 border-blue-200 pl-4 py-2">
                    <p className="text-lg font-medium text-blue-600 mb-1">{example.arabic}</p>
                    <p className="text-gray-700 mb-1">{example.english}</p>
                    <p className="text-sm text-gray-500 italic">/{example.pronunciation}/</p>
                  </div>
                ))}
              </div>
            </CardContent>
          </Card>

          {/* Related Words */}
          <Card>
            <CardHeader>
              <CardTitle>Related Words</CardTitle>
              <CardDescription>Expand your vocabulary with related terms</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                {word.relatedWords.map((relatedWord, index) => (
                  <div key={index} className="p-3 border rounded-lg hover:bg-gray-50 transition-colors">
                    <p className="font-medium text-blue-600">{relatedWord.arabic}</p>
                    <p className="text-gray-700">{relatedWord.english}</p>
                    <p className="text-sm text-gray-500 italic">/{relatedWord.pronunciation}/</p>
                  </div>
                ))}
              </div>
            </CardContent>
          </Card>
        </div>

        <div className="space-y-6">
          {/* Actions */}
          <Card>
            <CardHeader>
              <CardTitle>Actions</CardTitle>
            </CardHeader>
            <CardContent className="space-y-3">
              <Button className="w-full">
                <BookOpen className="mr-2 h-4 w-4" />
                Practice This Word
              </Button>
              <Button variant="outline" className="w-full">
                Add to Study Group
              </Button>
              <Button variant="outline" className="w-full">
                {word.learned ? "Mark as Unlearned" : "Mark as Learned"}
              </Button>
            </CardContent>
          </Card>

          {/* Word Groups */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center space-x-2">
                <Users className="h-5 w-5" />
                <span>Word Groups</span>
              </CardTitle>
              <CardDescription>This word belongs to these groups</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-2">
                {word.groups.map((group, index) => (
                  <Link key={index} to="/groups">
                    <Badge variant="outline" className="block text-center py-2 hover:bg-gray-50 cursor-pointer">
                      {group}
                    </Badge>
                  </Link>
                ))}
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
};

export default WordDetail;
