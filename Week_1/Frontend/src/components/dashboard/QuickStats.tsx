
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { TrendingUp, Calendar, Users, Flame } from "lucide-react";

export const QuickStats = () => {
  // Mock data - in a real app, this would come from GET /api/dashboard/quick_stats
  const quickStats = {
    successRate: 82, // percentage
    totalStudySessions: 47,
    totalActiveGroups: 8,
    studyStreak: 12, // days
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle>Quick Stats</CardTitle>
        <CardDescription>Your learning metrics at a glance</CardDescription>
      </CardHeader>
      <CardContent>
        <div className="grid grid-cols-2 gap-4">
          <div className="text-center p-3 bg-green-50 rounded-lg">
            <div className="flex items-center justify-center mb-2">
              <TrendingUp className="h-5 w-5 text-green-600" />
            </div>
            <p className="text-2xl font-bold text-green-600">{quickStats.successRate}%</p>
            <p className="text-xs text-gray-600">Success Rate</p>
          </div>

          <div className="text-center p-3 bg-blue-50 rounded-lg">
            <div className="flex items-center justify-center mb-2">
              <Calendar className="h-5 w-5 text-blue-600" />
            </div>
            <p className="text-2xl font-bold text-blue-600">{quickStats.totalStudySessions}</p>
            <p className="text-xs text-gray-600">Study Sessions</p>
          </div>

          <div className="text-center p-3 bg-purple-50 rounded-lg">
            <div className="flex items-center justify-center mb-2">
              <Users className="h-5 w-5 text-purple-600" />
            </div>
            <p className="text-2xl font-bold text-purple-600">{quickStats.totalActiveGroups}</p>
            <p className="text-xs text-gray-600">Active Groups</p>
          </div>

          <div className="text-center p-3 bg-orange-50 rounded-lg">
            <div className="flex items-center justify-center mb-2">
              <Flame className="h-5 w-5 text-orange-600" />
            </div>
            <p className="text-2xl font-bold text-orange-600">{quickStats.studyStreak}</p>
            <p className="text-xs text-gray-600">Day Streak</p>
          </div>
        </div>
      </CardContent>
    </Card>
  );
};
