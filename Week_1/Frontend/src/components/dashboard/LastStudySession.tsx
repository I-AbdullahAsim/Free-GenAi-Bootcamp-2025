
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Calendar, Clock, BookOpen } from "lucide-react";
import { Link } from "react-router-dom";

export const LastStudySession = () => {
  // Mock data - in a real app, this would come from GET /api/dashboard/last_study_session
  const lastSession = {
    id: 1,
    activityName: "Daily Vocabulary Review",
    groupName: "Basic Vocabulary",
    groupId: 1,
    completedAt: "2024-12-17 2:30 PM",
    correct: 18,
    incorrect: 7,
    totalWords: 25,
  };

  const successRate = Math.round((lastSession.correct / lastSession.totalWords) * 100);

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center space-x-2">
          <Clock className="h-5 w-5" />
          <span>Last Study Session</span>
        </CardTitle>
        <CardDescription>Your most recent learning activity</CardDescription>
      </CardHeader>
      <CardContent className="space-y-4">
        <div>
          <h3 className="font-semibold text-lg">{lastSession.activityName}</h3>
          <div className="flex items-center space-x-2 text-sm text-gray-600 mt-1">
            <Calendar className="h-4 w-4" />
            <span>{lastSession.completedAt}</span>
          </div>
        </div>

        <div className="space-y-2">
          <div className="flex items-center justify-between">
            <span className="text-sm text-gray-600">Group:</span>
            <Link to={`/groups/${lastSession.groupId}`}>
              <Badge variant="outline" className="hover:bg-gray-50 cursor-pointer">
                {lastSession.groupName}
              </Badge>
            </Link>
          </div>
          
          <div className="grid grid-cols-3 gap-4 text-center">
            <div>
              <p className="text-2xl font-bold text-green-600">{lastSession.correct}</p>
              <p className="text-xs text-gray-500">Correct</p>
            </div>
            <div>
              <p className="text-2xl font-bold text-red-600">{lastSession.incorrect}</p>
              <p className="text-xs text-gray-500">Incorrect</p>
            </div>
            <div>
              <p className="text-2xl font-bold text-blue-600">{successRate}%</p>
              <p className="text-xs text-gray-500">Success</p>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
};
