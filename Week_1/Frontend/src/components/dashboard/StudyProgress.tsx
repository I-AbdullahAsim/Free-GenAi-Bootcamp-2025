
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";
import { BookOpen, Target } from "lucide-react";

export const StudyProgress = () => {
  // Mock data - in a real app, this would come from GET /api/dashboard/study_progress
  const studyProgress = {
    totalWordsStudied: 87,
    totalWordsAvailable: 324,
    masteryProgress: 27, // percentage
  };

  const studyPercentage = Math.round((studyProgress.totalWordsStudied / studyProgress.totalWordsAvailable) * 100);

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center space-x-2">
          <Target className="h-5 w-5" />
          <span>Study Progress</span>
        </CardTitle>
        <CardDescription>Your overall learning journey</CardDescription>
      </CardHeader>
      <CardContent className="space-y-6">
        <div className="space-y-2">
          <div className="flex items-center justify-between">
            <span className="text-sm text-gray-600">Words Studied</span>
            <span className="font-semibold">
              {studyProgress.totalWordsStudied}/{studyProgress.totalWordsAvailable}
            </span>
          </div>
          <Progress value={studyPercentage} className="h-3" />
          <p className="text-xs text-gray-500 text-center">{studyPercentage}% of total vocabulary</p>
        </div>

        <div className="space-y-2">
          <div className="flex items-center justify-between">
            <span className="text-sm text-gray-600">Mastery Progress</span>
            <span className="font-semibold">{studyProgress.masteryProgress}%</span>
          </div>
          <Progress value={studyProgress.masteryProgress} className="h-3" />
          <p className="text-xs text-gray-500 text-center">
            Based on retention and accuracy across all sessions
          </p>
        </div>
      </CardContent>
    </Card>
  );
};
