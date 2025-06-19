
import { Link, useLocation } from "react-router-dom";
import { cn } from "@/lib/utils";
import { BookOpen, LayoutDashboard, List, Settings, Calendar, Users, Languages } from "lucide-react";
import { MobileSidebar } from "./MobileSidebar";

const navigationItems = [{
  name: "Dashboard",
  href: "/dashboard",
  icon: LayoutDashboard
}, {
  name: "Study Activities",
  href: "/study-activities",
  icon: BookOpen
}, {
  name: "Words",
  href: "/words",
  icon: List
}, {
  name: "Word Groups",
  href: "/groups",
  icon: Users
}, {
  name: "Sessions",
  href: "/sessions",
  icon: Calendar
}, {
  name: "Arabic Letters",
  href: "/arabic-letters",
  icon: Languages
}, {
  name: "Settings",
  href: "/settings",
  icon: Settings
}];

export const Navigation = () => {
  const location = useLocation();
  
  return (
    <>
      {/* Mobile Navigation Bar */}
      <nav className="md:hidden bg-white/80 dark:bg-slate-900/80 backdrop-blur-md border-b border-blue-100 dark:border-slate-700 sticky top-0 z-50 transition-colors duration-300">
        <div className="container mx-auto px-6">
          <div className="flex items-center justify-between h-16">
            <div className="flex items-center space-x-4">
              <MobileSidebar />
              <Link to="/dashboard" className="flex items-center space-x-2">
                <div className="w-8 h-8 bg-gradient-to-br from-blue-600 to-purple-600 rounded-lg flex items-center justify-center">
                  <span className="text-white font-bold text-sm">ع</span>
                </div>
                <span className="text-xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">Lang Portal</span>
              </Link>
            </div>
          </div>
        </div>
      </nav>

      {/* Desktop Right Sidebar */}
      <nav className="hidden md:block fixed right-0 top-0 h-full w-64 bg-white/95 dark:bg-slate-900/95 backdrop-blur-md border-l border-blue-100 dark:border-slate-700 z-40 transition-colors duration-300">
        <div className="flex flex-col h-full">
          {/* Header */}
          <div className="p-6 border-b border-gray-200 dark:border-slate-700">
            <Link to="/dashboard" className="flex items-center space-x-2">
              <div className="w-8 h-8 bg-gradient-to-br from-blue-600 to-purple-600 rounded-lg flex items-center justify-center">
                <span className="text-white font-bold text-sm">ع</span>
              </div>
              <span className="text-xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">Lang Portal</span>
            </Link>
          </div>
          
          {/* Navigation Items */}
          <div className="flex-1 p-4">
            <div className="space-y-2">
              {navigationItems.map(item => {
                const isActive = location.pathname === item.href || (item.href !== "/dashboard" && location.pathname.startsWith(item.href));
                const Icon = item.icon;
                return (
                  <Link
                    key={item.name}
                    to={item.href}
                    className={cn(
                      "flex items-center space-x-3 px-4 py-3 rounded-lg transition-all duration-200 text-sm font-medium w-full",
                      isActive
                        ? "bg-blue-100 dark:bg-blue-900/50 text-blue-700 dark:text-blue-300 shadow-sm"
                        : "text-gray-600 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400 hover:bg-blue-50 dark:hover:bg-slate-800"
                    )}
                  >
                    <Icon size={20} />
                    <span>{item.name}</span>
                  </Link>
                );
              })}
            </div>
          </div>
        </div>
      </nav>
    </>
  );
};
