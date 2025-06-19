
import { useState } from "react";
import { Link, useLocation } from "react-router-dom";
import { cn } from "@/lib/utils";
import { BookOpen, LayoutDashboard, List, Settings, Calendar, Users, Languages, Menu, X } from "lucide-react";
import { Button } from "@/components/ui/button";
import {
  Sheet,
  SheetContent,
  SheetTrigger,
} from "@/components/ui/sheet";

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

export const MobileSidebar = () => {
  const location = useLocation();
  const [open, setOpen] = useState(false);
  
  return (
    <Sheet open={open} onOpenChange={setOpen}>
      <SheetTrigger asChild>
        <Button variant="ghost" size="icon" className="md:hidden">
          <Menu className="h-6 w-6" />
        </Button>
      </SheetTrigger>
      <SheetContent side="left" className="w-64 p-0">
        <div className="flex flex-col h-full bg-white dark:bg-slate-900">
          <div className="p-6 border-b border-gray-200 dark:border-slate-700">
            <div className="flex items-center justify-between">
              <Link 
                to="/dashboard" 
                className="flex items-center space-x-2"
                onClick={() => setOpen(false)}
              >
                <div className="w-8 h-8 bg-gradient-to-br from-blue-600 to-purple-600 rounded-lg flex items-center justify-center">
                  <span className="text-white font-bold text-sm">ع</span>
                </div>
                <span className="text-xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">Lang Portal</span>
              </Link>
            </div>
          </div>
          
          <nav className="flex-1 p-4">
            <div className="space-y-2">
              {navigationItems.map(item => {
                const isActive = location.pathname === item.href || (item.href !== "/dashboard" && location.pathname.startsWith(item.href));
                const Icon = item.icon;
                return (
                  <Link
                    key={item.name}
                    to={item.href}
                    onClick={() => setOpen(false)}
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
          </nav>
        </div>
      </SheetContent>
    </Sheet>
  );
};
