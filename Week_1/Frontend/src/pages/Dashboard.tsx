
import { Button } from "@/components/ui/button";
import { Play } from "lucide-react";
import { Link } from "react-router-dom";
import { LastStudySession } from "@/components/dashboard/LastStudySession";
import { StudyProgress } from "@/components/dashboard/StudyProgress";
import { QuickStats } from "@/components/dashboard/QuickStats";

const Dashboard = () => {
  return (
    <div className="space-y-8">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-gray-900 mb-2">Dashboard</h1>
          <p className="text-gray-600">Welcome back! Continue your Arabic learning journey.</p>
        </div>
        <Link to="/study-activities">
          <Button size="lg" className="bg-blue-600 hover:bg-blue-700">
            <Play className="mr-2 h-5 w-5" />
            Start Studying
          </Button>
        </Link>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <div className="lg:col-span-2 space-y-6">
          <LastStudySession />
          <StudyProgress />
        </div>
        
        <div className="space-y-6">
          <QuickStats />
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
