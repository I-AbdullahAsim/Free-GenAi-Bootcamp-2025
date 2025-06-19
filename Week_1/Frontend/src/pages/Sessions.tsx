
import { useState } from "react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Calendar, Clock, BookOpen, TrendingUp, ChevronUp, ChevronDown } from "lucide-react";
import { Link } from "react-router-dom";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination";

type SortField = 'date' | 'activity' | 'groupName' | 'duration' | 'reviewItems' | 'score';
type SortDirection = 'asc' | 'desc';

const Sessions = () => {
  const [currentPage, setCurrentPage] = useState(1);
  const [sortField, setSortField] = useState<SortField>('date');
  const [sortDirection, setSortDirection] = useState<SortDirection>('desc');

  const sessions = [
    {
      id: 1,
      date: "2024-12-17",
      startTime: "2:30 PM",
      endTime: "2:50 PM",
      activity: "Daily Vocabulary Review",
      groupName: "Core Words",
      groupId: 1,
      reviewItems: 25,
      duration: 20,
      score: 92,
    },
    {
      id: 2,
      date: "2024-12-16",
      startTime: "1:15 PM", 
      endTime: "1:45 PM",
      activity: "Grammar Fundamentals",
      groupName: "Basic Grammar",
      groupId: 2,
      reviewItems: 18,
      duration: 30,
      score: 87,
    },
    // Add more sessions for pagination demo
    ...Array.from({ length: 98 }, (_, i) => ({
      id: i + 3,
      date: new Date(2024, 11, 15 - i).toISOString().split('T')[0],
      startTime: `${Math.floor(Math.random() * 12) + 1}:${String(Math.floor(Math.random() * 60)).padStart(2, '0')} ${Math.random() > 0.5 ? 'PM' : 'AM'}`,
      endTime: `${Math.floor(Math.random() * 12) + 1}:${String(Math.floor(Math.random() * 60)).padStart(2, '0')} ${Math.random() > 0.5 ? 'PM' : 'AM'}`,
      activity: ["Daily Vocabulary Review", "Grammar Fundamentals", "Conversation Practice", "Reading Comprehension"][Math.floor(Math.random() * 4)],
      groupName: ["Core Words", "Basic Grammar", "Advanced Vocabulary", "Conversation"][Math.floor(Math.random() * 4)],
      groupId: Math.floor(Math.random() * 4) + 1,
      reviewItems: Math.floor(Math.random() * 50) + 10,
      duration: Math.floor(Math.random() * 45) + 15,
      score: Math.floor(Math.random() * 30) + 70,
    }))
  ];

  const sessionsPerPage = 50;
  const totalPages = Math.ceil(sessions.length / sessionsPerPage);

  const sortedSessions = [...sessions].sort((a, b) => {
    let aValue: any = a[sortField];
    let bValue: any = b[sortField];
    
    if (sortField === 'date') {
      aValue = new Date(aValue).getTime();
      bValue = new Date(bValue).getTime();
    }
    
    if (typeof aValue === 'number' && typeof bValue === 'number') {
      return sortDirection === 'asc' ? aValue - bValue : bValue - aValue;
    }
    
    const comparison = aValue.toString().localeCompare(bValue.toString());
    return sortDirection === 'asc' ? comparison : -comparison;
  });

  const paginatedSessions = sortedSessions.slice(
    (currentPage - 1) * sessionsPerPage,
    currentPage * sessionsPerPage
  );

  const handleSort = (field: SortField) => {
    if (sortField === field) {
      setSortDirection(sortDirection === 'asc' ? 'desc' : 'asc');
    } else {
      setSortField(field);
      setSortDirection('asc');
    }
  };

  const getSortIcon = (field: SortField) => {
    if (sortField !== field) return null;
    return sortDirection === 'asc' ? 
      <ChevronUp className="h-4 w-4 inline ml-1" /> : 
      <ChevronDown className="h-4 w-4 inline ml-1" />;
  };

  const getScoreColor = (score: number) => {
    if (score >= 90) return "text-green-600 bg-green-100";
    if (score >= 80) return "text-yellow-600 bg-yellow-100";
    if (score >= 70) return "text-orange-600 bg-orange-100";
    return "text-red-600 bg-red-100";
  };

  const totalSessions = sessions.length;
  const totalReviewItems = sessions.reduce((sum, session) => sum + session.reviewItems, 0);
  const averageScore = Math.round(sessions.reduce((sum, session) => sum + session.score, 0) / sessions.length);
  const totalDuration = sessions.reduce((sum, session) => sum + session.duration, 0);

  return (
    <div className="space-y-8">
      <div>
        <h1 className="text-3xl font-bold text-gray-900 mb-2">Study Sessions</h1>
        <p className="text-gray-600">Track your learning progress and study history.</p>
      </div>

      {/* Summary Stats */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">Total Sessions</p>
                <p className="text-2xl font-bold">{totalSessions}</p>
              </div>
              <Calendar className="h-8 w-8 text-blue-600" />
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">Review Items</p>
                <p className="text-2xl font-bold">{totalReviewItems}</p>
              </div>
              <BookOpen className="h-8 w-8 text-green-600" />
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">Average Score</p>
                <p className="text-2xl font-bold">{averageScore}%</p>
              </div>
              <TrendingUp className="h-8 w-8 text-purple-600" />
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">Total Time</p>
                <p className="text-2xl font-bold">{Math.floor(totalDuration / 60)}h {totalDuration % 60}m</p>
              </div>
              <Clock className="h-8 w-8 text-orange-600" />
            </div>
          </CardContent>
        </Card>
      </div>

      {/* Sessions Table */}
      <Card>
        <CardHeader>
          <CardTitle>Session History</CardTitle>
          <CardDescription>Your complete study session log</CardDescription>
        </CardHeader>
        <CardContent>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>
                  <button 
                    onClick={() => handleSort('date')}
                    className="flex items-center hover:text-blue-600 font-medium"
                  >
                    Date {getSortIcon('date')}
                  </button>
                </TableHead>
                <TableHead>
                  <button 
                    onClick={() => handleSort('activity')}
                    className="flex items-center hover:text-blue-600 font-medium"
                  >
                    Activity {getSortIcon('activity')}
                  </button>
                </TableHead>
                <TableHead>
                  <button 
                    onClick={() => handleSort('groupName')}
                    className="flex items-center hover:text-blue-600 font-medium"
                  >
                    Group {getSortIcon('groupName')}
                  </button>
                </TableHead>
                <TableHead>Time</TableHead>
                <TableHead>
                  <button 
                    onClick={() => handleSort('duration')}
                    className="flex items-center hover:text-blue-600 font-medium"
                  >
                    Duration {getSortIcon('duration')}
                  </button>
                </TableHead>
                <TableHead>
                  <button 
                    onClick={() => handleSort('reviewItems')}
                    className="flex items-center hover:text-blue-600 font-medium"
                  >
                    Review Items {getSortIcon('reviewItems')}
                  </button>
                </TableHead>
                <TableHead>
                  <button 
                    onClick={() => handleSort('score')}
                    className="flex items-center hover:text-blue-600 font-medium"
                  >
                    Score {getSortIcon('score')}
                  </button>
                </TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {paginatedSessions.map((session) => (
                <TableRow key={session.id}>
                  <TableCell className="font-medium">{session.date}</TableCell>
                  <TableCell>{session.activity}</TableCell>
                  <TableCell>
                    <Link 
                      to={`/groups/${session.groupId}`}
                      className="text-blue-600 hover:text-blue-800 hover:underline"
                    >
                      {session.groupName}
                    </Link>
                  </TableCell>
                  <TableCell className="text-sm text-gray-600">
                    {session.startTime} - {session.endTime}
                  </TableCell>
                  <TableCell>{session.duration} min</TableCell>
                  <TableCell>{session.reviewItems}</TableCell>
                  <TableCell>
                    <Badge className={getScoreColor(session.score)}>
                      {session.score}%
                    </Badge>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </CardContent>
      </Card>

      {/* Pagination */}
      <Pagination>
        <PaginationContent>
          <PaginationItem>
            <PaginationPrevious 
              onClick={() => setCurrentPage(Math.max(1, currentPage - 1))}
              className={currentPage === 1 ? "pointer-events-none opacity-50" : "cursor-pointer"}
            />
          </PaginationItem>
          
          {Array.from({ length: Math.min(5, totalPages) }, (_, i) => {
            const page = Math.max(1, Math.min(totalPages - 4, currentPage - 2)) + i;
            return (
              <PaginationItem key={page}>
                <PaginationLink
                  onClick={() => setCurrentPage(page)}
                  isActive={page === currentPage}
                  className="cursor-pointer"
                >
                  {page}
                </PaginationLink>
              </PaginationItem>
            );
          })}
          
          <PaginationItem>
            <PaginationNext 
              onClick={() => setCurrentPage(Math.min(totalPages, currentPage + 1))}
              className={currentPage === totalPages ? "pointer-events-none opacity-50" : "cursor-pointer"}
            />
          </PaginationItem>
        </PaginationContent>
      </Pagination>
    </div>
  );
};

export default Sessions;
