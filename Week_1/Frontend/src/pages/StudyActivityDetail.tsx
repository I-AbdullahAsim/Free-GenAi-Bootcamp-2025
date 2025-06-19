
import { useParams } from "react-router-dom";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { ArrowLeft, Clock, BookOpen, Play } from "lucide-react";
import { Link } from "react-router-dom";

const StudyActivityDetail = () => {
  const { id } = useParams();

  // Mock data - in a real app, this would be fetched based on the ID
  const activity = {
    id: parseInt(id || "1"),
    title: "Daily Vocabulary Review",
    description: "Review and practice your learned vocabulary words with interactive exercises and quizzes.",
    duration: "15-20 min",
    difficulty: "Beginner",
    wordsCount: 25,
    lastCompleted: "2 hours ago",
    nextWords: ["بيت", "كتاب", "مدرسة", "طعام", "ماء"],
    objectives: [
      "Review previously learned vocabulary",
      "Practice pronunciation",
      "Complete comprehension exercises",
      "Test retention with quizzes"
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

  return (
    <div className="space-y-8">
      <div className="flex items-center space-x-4">
        <Link to="/study-activities">
          <Button variant="outline" size="sm">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back to Activities
          </Button>
        </Link>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <div className="lg:col-span-2 space-y-6">
          <div>
            <div className="flex items-start justify-between mb-4">
              <div>
                <h1 className="text-3xl font-bold text-gray-900 mb-2">{activity.title}</h1>
                <p className="text-gray-600">{activity.description}</p>
              </div>
              <Badge className={getDifficultyColor(activity.difficulty)}>
                {activity.difficulty}
              </Badge>
            </div>

            <div className="flex items-center space-x-6 text-sm text-gray-600">
              <div className="flex items-center space-x-1">
                <Clock className="h-4 w-4" />
                <span>{activity.duration}</span>
              </div>
              <div className="flex items-center space-x-1">
                <BookOpen className="h-4 w-4" />
                <span>{activity.wordsCount} words</span>
              </div>
            </div>
          </div>

          <Card>
            <CardHeader>
              <CardTitle>Learning Objectives</CardTitle>
              <CardDescription>What you'll accomplish in this session</CardDescription>
            </CardHeader>
            <CardContent>
              <ul className="space-y-2">
                {activity.objectives.map((objective, index) => (
                  <li key={index} className="flex items-start space-x-2">
                    <span className="w-2 h-2 bg-blue-600 rounded-full mt-2 flex-shrink-0"></span>
                    <span>{objective}</span>
                  </li>
                ))}
              </ul>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Words in This Session</CardTitle>
              <CardDescription>Preview of vocabulary you'll be practicing</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="flex flex-wrap gap-2">
                {activity.nextWords.map((word, index) => (
                  <Badge key={index} variant="outline" className="text-lg py-2 px-3">
                    {word}
                  </Badge>
                ))}
              </div>
            </CardContent>
          </Card>
        </div>

        <div className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle>Ready to Start?</CardTitle>
              <CardDescription>Begin your learning session</CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              <Button className="w-full" size="lg">
                <Play className="mr-2 h-5 w-5" />
                Start Session
              </Button>
              <p className="text-sm text-gray-500 text-center">
                Last completed: {activity.lastCompleted}
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Session Tips</CardTitle>
            </CardHeader>
            <CardContent className="space-y-3 text-sm">
              <p>• Find a quiet environment for better focus</p>
              <p>• Take breaks if you feel overwhelmed</p>
              <p>• Practice pronunciation out loud</p>
              <p>• Review difficult words multiple times</p>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
};

export default StudyActivityDetail;
