
import { useLocation, Link } from "react-router-dom";
import { ChevronRight, Home } from "lucide-react";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";

export const Breadcrumbs = () => {
  const location = useLocation();
  const pathSegments = location.pathname.split('/').filter(Boolean);
  
  const getBreadcrumbText = (segment: string, index: number) => {
    switch (segment) {
      case 'dashboard': return 'Dashboard';
      case 'study-activities': return 'Study Activities';
      case 'words': return 'Words';
      case 'groups': return 'Word Groups';
      case 'sessions': return 'Sessions';
      case 'arabic-letters': return 'Arabic Letters';
      case 'settings': return 'Settings';
      default:
        // For dynamic segments like IDs, show placeholder names
        if (pathSegments[index - 1] === 'study-activities') return 'Adventure MUD';
        if (pathSegments[index - 1] === 'words') return 'مرحباً';
        if (pathSegments[index - 1] === 'groups') return 'Core Words';
        return segment;
    }
  };

  const getBreadcrumbPath = (index: number) => {
    return '/' + pathSegments.slice(0, index + 1).join('/');
  };

  if (pathSegments.length === 0) return null;

  return (
    <div className="bg-white/90 dark:bg-slate-800/90 backdrop-blur-sm border-b border-gray-200 dark:border-slate-700 px-6 py-3">
      <Breadcrumb>
        <BreadcrumbList>
          <BreadcrumbItem>
            <BreadcrumbLink asChild>
              <Link to="/dashboard" className="flex items-center text-gray-600 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400">
                <Home className="h-4 w-4" />
              </Link>
            </BreadcrumbLink>
          </BreadcrumbItem>
          
          {pathSegments.map((segment, index) => (
            <div key={index} className="flex items-center">
              <BreadcrumbSeparator>
                <ChevronRight className="h-4 w-4 text-gray-400 dark:text-gray-500" />
              </BreadcrumbSeparator>
              <BreadcrumbItem>
                {index === pathSegments.length - 1 ? (
                  <BreadcrumbPage className="text-gray-900 dark:text-gray-100 font-medium">
                    {getBreadcrumbText(segment, index)}
                  </BreadcrumbPage>
                ) : (
                  <BreadcrumbLink asChild>
                    <Link to={getBreadcrumbPath(index)} className="text-gray-600 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400">
                      {getBreadcrumbText(segment, index)}
                    </Link>
                  </BreadcrumbLink>
                )}
              </BreadcrumbItem>
            </div>
          ))}
        </BreadcrumbList>
      </Breadcrumb>
    </div>
  );
};
