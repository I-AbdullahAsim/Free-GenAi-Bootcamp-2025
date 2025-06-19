
import { useState } from "react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import { Search, BookOpen, Filter, Volume2, ChevronUp, ChevronDown } from "lucide-react";
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

type SortField = 'arabic' | 'english' | 'correct' | 'incorrect';
type SortDirection = 'asc' | 'desc';

const Words = () => {
  const [searchTerm, setSearchTerm] = useState("");
  const [currentPage, setCurrentPage] = useState(1);
  const [sortField, setSortField] = useState<SortField>('arabic');
  const [sortDirection, setSortDirection] = useState<SortDirection>('asc');

  const words = [
    { id: 1, arabic: "بيت", english: "house", pronunciation: "bayt", correct: 15, incorrect: 3 },
    { id: 2, arabic: "كتاب", english: "book", pronunciation: "kitab", correct: 12, incorrect: 1 },
    { id: 3, arabic: "مدرسة", english: "school", pronunciation: "madrasa", correct: 8, incorrect: 5 },
    { id: 4, arabic: "طعام", english: "food", pronunciation: "ta'am", correct: 20, incorrect: 2 },
    { id: 5, arabic: "سيارة", english: "car", pronunciation: "sayyara", correct: 6, incorrect: 4 },
    { id: 6, arabic: "مستشفى", english: "hospital", pronunciation: "mustashfa", correct: 3, incorrect: 7 },
    // Add more words to demonstrate pagination
    ...Array.from({ length: 94 }, (_, i) => ({
      id: i + 7,
      arabic: `كلمة${i + 7}`,
      english: `word${i + 7}`,
      pronunciation: `word${i + 7}`,
      correct: Math.floor(Math.random() * 25),
      incorrect: Math.floor(Math.random() * 10),
    }))
  ];

  const wordsPerPage = 50;
  const totalPages = Math.ceil(words.length / wordsPerPage);

  const filteredWords = words.filter((word) =>
    word.arabic.includes(searchTerm) ||
    word.english.toLowerCase().includes(searchTerm.toLowerCase()) ||
    word.pronunciation.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const sortedWords = [...filteredWords].sort((a, b) => {
    const aValue = a[sortField];
    const bValue = b[sortField];
    
    if (typeof aValue === 'number' && typeof bValue === 'number') {
      return sortDirection === 'asc' ? aValue - bValue : bValue - aValue;
    }
    
    const comparison = aValue.toString().localeCompare(bValue.toString());
    return sortDirection === 'asc' ? comparison : -comparison;
  });

  const paginatedWords = sortedWords.slice(
    (currentPage - 1) * wordsPerPage,
    currentPage * wordsPerPage
  );

  const handleSort = (field: SortField) => {
    if (sortField === field) {
      setSortDirection(sortDirection === 'asc' ? 'desc' : 'asc');
    } else {
      setSortField(field);
      setSortDirection('asc');
    }
  };

  const handlePlayAudio = (pronunciation: string) => {
    // Placeholder for audio playback
    console.log(`Playing audio for: ${pronunciation}`);
    // In a real app, this would use Web Speech API or audio files
  };

  const getSortIcon = (field: SortField) => {
    if (sortField !== field) return null;
    return sortDirection === 'asc' ? 
      <ChevronUp className="h-4 w-4 inline ml-1" /> : 
      <ChevronDown className="h-4 w-4 inline ml-1" />;
  };

  return (
    <div className="space-y-8">
      <div>
        <h1 className="text-3xl font-bold text-gray-900 mb-2">Vocabulary</h1>
        <p className="text-gray-600">Explore and manage your Arabic vocabulary collection.</p>
      </div>

      {/* Search and Filter */}
      <div className="flex items-center space-x-4">
        <div className="relative flex-1 max-w-md">
          <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 h-4 w-4" />
          <Input
            placeholder="Search words..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            className="pl-10"
          />
        </div>
        <Button variant="outline">
          <Filter className="mr-2 h-4 w-4" />
          Filter
        </Button>
      </div>

      {/* Stats */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">Total Words</p>
                <p className="text-2xl font-bold">{words.length}</p>
              </div>
              <BookOpen className="h-8 w-8 text-blue-600" />
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">Average Correct</p>
                <p className="text-2xl font-bold">{Math.round(words.reduce((sum, w) => sum + w.correct, 0) / words.length)}</p>
              </div>
              <div className="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                <span className="text-green-600 font-bold">✓</span>
              </div>
            </div>
          </CardContent>
        </Card>
        <Card>
          <CardContent className="pt-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">Average Incorrect</p>
                <p className="text-2xl font-bold">{Math.round(words.reduce((sum, w) => sum + w.incorrect, 0) / words.length)}</p>
              </div>
              <div className="w-8 h-8 bg-red-100 rounded-full flex items-center justify-center">
                <span className="text-red-600 font-bold">✗</span>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>

      {/* Words Table */}
      <Card>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>
                <button 
                  onClick={() => handleSort('arabic')}
                  className="flex items-center hover:text-blue-600 font-medium"
                >
                  Arabic {getSortIcon('arabic')}
                </button>
              </TableHead>
              <TableHead>
                <button 
                  onClick={() => handleSort('english')}
                  className="flex items-center hover:text-blue-600 font-medium"
                >
                  English {getSortIcon('english')}
                </button>
              </TableHead>
              <TableHead>
                <button 
                  onClick={() => handleSort('correct')}
                  className="flex items-center hover:text-blue-600 font-medium"
                >
                  Correct {getSortIcon('correct')}
                </button>
              </TableHead>
              <TableHead>
                <button 
                  onClick={() => handleSort('incorrect')}
                  className="flex items-center hover:text-blue-600 font-medium"
                >
                  Incorrect {getSortIcon('incorrect')}
                </button>
              </TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {paginatedWords.map((word) => (
              <TableRow key={word.id}>
                <TableCell>
                  <div className="flex items-center space-x-2">
                    <Link 
                      to={`/words/${word.id}`}
                      className="text-xl font-bold text-blue-600 hover:text-blue-800"
                    >
                      {word.arabic}
                    </Link>
                    <Button
                      variant="ghost"
                      size="sm"
                      onClick={() => handlePlayAudio(word.pronunciation)}
                      className="p-1 h-8 w-8"
                    >
                      <Volume2 className="h-4 w-4" />
                    </Button>
                  </div>
                </TableCell>
                <TableCell className="font-medium">{word.english}</TableCell>
                <TableCell>
                  <Badge className="bg-green-100 text-green-700">{word.correct}</Badge>
                </TableCell>
                <TableCell>
                  <Badge className="bg-red-100 text-red-700">{word.incorrect}</Badge>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
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
          
          {Array.from({ length: totalPages }, (_, i) => i + 1).map((page) => (
            <PaginationItem key={page}>
              <PaginationLink
                onClick={() => setCurrentPage(page)}
                isActive={page === currentPage}
                className="cursor-pointer"
              >
                {page}
              </PaginationLink>
            </PaginationItem>
          ))}
          
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

export default Words;
