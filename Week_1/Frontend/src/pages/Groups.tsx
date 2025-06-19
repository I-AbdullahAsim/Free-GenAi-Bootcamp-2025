
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Users, BookOpen, Plus } from "lucide-react";
import { Link } from "react-router-dom";

const Groups = () => {
  const groups = [
    {
      id: 1,
      name: "Basic Vocabulary",
      description: "Essential words for beginners",
      wordCount: 50,
      difficulty: "Beginner",
      lastStudied: "2 hours ago",
      progress: 85,
    },
    {
      id: 2,
      name: "Family & Home",
      description: "Words related to family and household items",
      wordCount: 30,
      difficulty: "Beginner",
      lastStudied: "1 day ago",
      progress: 60,
    },
    {
      id: 3,
      name: "Food & Drink",
      description: "Culinary vocabulary and dining terms",
      wordCount: 45,
      difficulty: "Intermediate",
      lastStudied: "3 days ago",
      progress: 40,
    },
    {
      id: 4,
      name: "Transportation",
      description: "Vehicles and travel-related vocabulary",
      wordCount: 25,
      difficulty: "Intermediate",
      lastStudied: "5 days ago",
      progress: 20,
    },
    {
      id: 5,
      name: "Business Arabic",
      description: "Professional and business terminology",
      wordCount: 75,
      difficulty: "Advanced",
      lastStudied: "1 week ago",
      progress: 15,
    },
    {
      id: 6,
      name: "Medical Terms",
      description: "Healthcare and medical vocabulary",
      wordCount: 60,
      difficulty: "Advanced",
      lastStudied: "2 weeks ago",
      progress: 5,
    },
  ];

  const getDifficultyColor = (difficulty: string) => {
    switch (difficulty) {
      case "Beginner": return "bg-green-100 text-green-700";
      case "Intermediate": return "bg-yellow-100 text-yellow-700";
      case "Advanced": return "bg-red-100 text-red-700";
      default: return "bg-gray-100 text-gray-700";
    }
  };

  const getProgressColor = (progress: number) => {
    if (progress >= 70) return "bg-green-500";
    if (progress >= 40) return "bg-yellow-500";
    return "bg-red-500";
  };

  return (
    <div className="space-y-8">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-gray-900 mb-2">Word Groups</h1>
          <p className="text-gray-600">Organize your vocabulary into themed collections.</p>
        </div>
        <Button>
          <Plus className="mr-2 h-4 w-4" />
          Create Group
        </Button>
      </div>

      {/* Summary Stats */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">Total Groups</p>
                <p className="text-2xl font-bold">{groups.length}</p>
              </div>
              <Users className="h-8 w-8 text-blue-600" />
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">Total Words</p>
                <p className="text-2xl font-bold">{groups.reduce((sum, group) => sum + group.wordCount, 0)}</p>
              </div>
              <BookOpen className="h-8 w-8 text-green-600" />
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">Active Groups</p>
                <p className="text-2xl font-bold">{groups.filter(g => g.progress > 0).length}</p>
              </div>
              <div className="w-8 h-8 bg-purple-100 rounded-full flex items-center justify-center">
                <span className="text-purple-600 font-bold">▶</span>
              </div>
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">Avg Progress</p>
                <p className="text-2xl font-bold">
                  {Math.round(groups.reduce((sum, group) => sum + group.progress, 0) / groups.length)}%
                </p>
              </div>
              <div className="w-8 h-8 bg-orange-100 rounded-full flex items-center justify-center">
                <span className="text-orange-600 font-bold">%</span>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>

      {/* Groups Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {groups.map((group) => (
          <Link key={group.id} to={`/groups/${group.id}`}>
            <Card className="hover:shadow-lg transition-shadow duration-200 cursor-pointer">
              <CardHeader>
                <div className="flex items-start justify-between">
                  <div>
                    <CardTitle className="text-lg">{group.name}</CardTitle>
                    <CardDescription>{group.description}</CardDescription>
                  </div>
                  <Badge className={getDifficultyColor(group.difficulty)}>
                    {group.difficulty}
                  </Badge>
                </div>
              </CardHeader>
              <CardContent className="space-y-4">
                <div className="flex items-center justify-between text-sm">
                  <span className="text-gray-600">{group.wordCount} words</span>
                  <span className="text-gray-500">Last studied: {group.lastStudied}</span>
                </div>
                
                <div className="space-y-2">
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-gray-600">Progress</span>
                    <span className="font-medium">{group.progress}%</span>
                  </div>
                  <div className="w-full bg-gray-200 rounded-full h-2">
                    <div
                      className={`h-2 rounded-full ${getProgressColor(group.progress)}`}
                      style={{ width: `${group.progress}%` }}
                    ></div>
                  </div>
                </div>
              </CardContent>
            </Card>
          </Link>
        ))}
      </div>
    </div>
  );
};

export default Groups;
