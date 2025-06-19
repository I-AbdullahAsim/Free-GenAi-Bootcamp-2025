
import { useParams } from "react-router-dom";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { ArrowLeft, BookOpen, Play, Plus, Settings } from "lucide-react";
import { Link } from "react-router-dom";

const GroupDetail = () => {
  const { id } = useParams();

  // Mock data - in a real app, this would be fetched based on the ID
  const group = {
    id: parseInt(id || "1"),
    name: "Basic Vocabulary",
    description: "Essential words for beginners learning Arabic",
    wordCount: 50,
    difficulty: "Beginner",
    lastStudied: "2 hours ago",
    progress: 85,
    words: [
      { id: 1, arabic: "بيت", english: "house", pronunciation: "bayt", learned: true },
      { id: 2, arabic: "كتاب", english: "book", pronunciation: "kitab", learned: true },
      { id: 3, arabic: "مدرسة", english: "school", pronunciation: "madrasa", learned: false },
      { id: 4, arabic: "طعام", english: "food", pronunciation: "ta'am", learned: true },
      { id: 5, arabic: "ماء", english: "water", pronunciation: "ma'", learned: false },
      { id: 6, arabic: "شمس", english: "sun", pronunciation: "shams", learned: true },
      { id: 7, arabic: "قمر", english: "moon", pronunciation: "qamar", learned: false },
      { id: 8, arabic: "نجم", english: "star", pronunciation: "najm", learned: true },
    ]
  };

  const getDifficultyColor = (difficulty: string) => {
    switch (difficulty) {
      case "Beginner": return "bg-green-100 text-green-700";
      case "Intermediate": return "bg-yellow-100 text-yellow-700";
      case "Advanced": return "bg-red-100 text-red-700";
      default: return "bg-gray-100 text-gray-700";
    }
  };

  const learnedCount = group.words.filter(word => word.learned).length;

  return (
    <div className="space-y-8">
      <div className="flex items-center space-x-4">
        <Link to="/groups">
          <Button variant="outline" size="sm">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Groups
          </Button>
        </Link>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <div className="lg:col-span-2 space-y-6">
          {/* Group Header */}
          <div>
            <div className="flex items-start justify-between mb-4">
              <div>
                <h1 className="text-3xl font-bold text-gray-900 mb-2">{group.name}</h1>
                <p className="text-gray-600">{group.description}</p>
              </div>
              <div className="flex items-center space-x-2">
                <Badge className={getDifficultyColor(group.difficulty)}>
                  {group.difficulty}
                </Badge>
                <Button variant="outline" size="sm">
                  <Settings className="h-4 w-4" />
                </Button>
              </div>
            </div>

            <div className="flex items-center space-x-6 text-sm text-gray-600">
              <span>{group.wordCount} words</span>
              <span>Last studied: {group.lastStudied}</span>
              <span>{learnedCount} learned</span>
            </div>

            {/* Progress Bar */}
            <div className="mt-4 space-y-2">
              <div className="flex items-center justify-between text-sm">
                <span className="text-gray-600">Progress</span>
                <span className="font-medium">{group.progress}%</span>
              </div>
              <div className="w-full bg-gray-200 rounded-full h-3">
                <div
                  className="h-3 rounded-full bg-green-500"
                  style={{ width: `${group.progress}%` }}
                ></div>
              </div>
            </div>
          </div>

          {/* Words List */}
          <Card>
            <CardHeader>
              <div className="flex items-center justify-between">
                <div>
                  <CardTitle>Words in this Group</CardTitle>
                  <CardDescription>Click on any word to view details</CardDescription>
                </div>
                <Button variant="outline" size="sm">
                  <Plus className="h-4 w-4 mr-2" />
                  Add Word
                </Button>
              </div>
            </CardHeader>
            <CardContent>
              <div className="space-y-3">
                {group.words.map((word) => (
                  <Link key={word.id} to={`/words/${word.id}`}>
                    <div className="flex items-center justify-between p-3 border rounded-lg hover:bg-gray-50 transition-colors cursor-pointer">
                      <div className="flex items-center space-x-4">
                        <div className="text-2xl font-bold text-blue-600">
                          {word.arabic}
                        </div>
                        <div>
                          <p className="font-medium">{word.english}</p>
                          <p className="text-sm text-gray-500 italic">/{word.pronunciation}/</p>
                        </div>
                      </div>
                      <div className="flex items-center space-x-2">
                        {word.learned && (
                          <div className="w-6 h-6 bg-green-100 rounded-full flex items-center justify-center">
                            <span className="text-green-600 text-sm font-bold">✓</span>
                          </div>
                        )}
                      </div>
                    </div>
                  </Link>
                ))}
              </div>
            </CardContent>
          </Card>
        </div>

        <div className="space-y-6">
          {/* Study Actions */}
          <Card>
            <CardHeader>
              <CardTitle>Study This Group</CardTitle>
              <CardDescription>Practice with focused learning sessions</CardDescription>
            </CardHeader>
            <CardContent className="space-y-3">
              <Button className="w-full" size="lg">
                <Play className="mr-2 h-5 w-5" />
                Start Study Session
              </Button>
              <Button variant="outline" className="w-full">
                <BookOpen className="mr-2 h-4 w-4" />
                Quick Review
              </Button>
              <Button variant="outline" className="w-full">
                Practice Unlearned
              </Button>
            </CardContent>
          </Card>

          {/* Group Stats */}
          <Card>
            <CardHeader>
              <CardTitle>Group Statistics</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="flex justify-between">
                <span className="text-gray-600">Total Words:</span>
                <span className="font-medium">{group.wordCount}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-600">Learned:</span>
                <span className="font-medium text-green-600">{learnedCount}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-600">To Learn:</span>
                <span className="font-medium text-orange-600">{group.wordCount - learnedCount}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-600">Difficulty:</span>
                <Badge className={getDifficultyColor(group.difficulty)} variant="outline">
                  {group.difficulty}
                </Badge>
              </div>
            </CardContent>
          </Card>

          {/* Study Tips */}
          <Card>
            <CardHeader>
              <CardTitle>Study Tips</CardTitle>
            </CardHeader>
            <CardContent className="space-y-3 text-sm">
              <p>• Focus on words you haven't learned yet</p>
              <p>• Practice pronunciation with each word</p>
              <p>• Create your own example sentences</p>
              <p>• Review learned words periodically</p>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
};

export default GroupDetail;
