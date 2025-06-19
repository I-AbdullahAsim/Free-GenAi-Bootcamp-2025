
import { Toaster } from "@/components/ui/toaster";
import { Toaster as Sonner } from "@/components/ui/sonner";
import { TooltipProvider } from "@/components/ui/tooltip";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import { Navigation } from "./components/Navigation";
import { Breadcrumbs } from "./components/Breadcrumbs";
import Dashboard from "./pages/Dashboard";
import StudyActivities from "./pages/StudyActivities";
import StudyActivityDetail from "./pages/StudyActivityDetail";
import Words from "./pages/Words";
import WordDetail from "./pages/WordDetail";
import Groups from "./pages/Groups";
import GroupDetail from "./pages/GroupDetail";
import Sessions from "./pages/Sessions";
import Settings from "./pages/Settings";
import ArabicLetters from "./pages/ArabicLetters";
import NotFound from "./pages/NotFound";

const queryClient = new QueryClient();

const App = () => (
  <QueryClientProvider client={queryClient}>
    <TooltipProvider>
      <Toaster />
      <Sonner />
      <BrowserRouter>
        <div className="min-h-screen bg-gradient-to-br from-slate-50 to-blue-50 dark:from-slate-900 dark:to-slate-800 transition-colors duration-300">
          <Navigation />
          <div className="md:mr-64">
            <Breadcrumbs />
            <main className="container mx-auto px-6 py-8">
              <Routes>
                <Route path="/" element={<Navigate to="/dashboard" replace />} />
                <Route path="/dashboard" element={<Dashboard />} />
                <Route path="/study-activities" element={<StudyActivities />} />
                <Route path="/study-activities/:id" element={<StudyActivityDetail />} />
                <Route path="/words" element={<Words />} />
                <Route path="/words/:id" element={<WordDetail />} />
                <Route path="/groups" element={<Groups />} />
                <Route path="/groups/:id" element={<GroupDetail />} />
                <Route path="/sessions" element={<Sessions />} />
                <Route path="/arabic-letters" element={<ArabicLetters />} />
                <Route path="/settings" element={<Settings />} />
                <Route path="*" element={<NotFound />} />
              </Routes>
            </main>
          </div>
        </div>
      </BrowserRouter>
    </TooltipProvider>
  </QueryClientProvider>
);

export default App;
