import { useState } from 'react';
import { Outlet, Link, useLocation } from 'react-router-dom';
import { Users, List, BarChart2, Settings, Menu, X, Bell, Search } from 'lucide-react';

export default function AdminPage() {
    const [collapsed, setCollapsed] = useState(false);
    const location = useLocation();

    // Navigation items
    const navItems = [
        { path: "/admin/handlers", name: "Handlers", icon: <Users size={20} /> },
        { path: "/admin/tasks", name: "Tasks", icon: <List size={20} /> },
        { path: "/admin/analysts", name: "Analysts", icon: <BarChart2 size={20} /> },
        { path: "/admin/settings", name: "Settings", icon: <Settings size={20} /> }
    ];

    // Check if current path matches the nav item path
    const isActive = (path: string) => {
        return location.pathname.startsWith(path);
    };

    return (
        <div className="flex h-screen bg-gray-100">
            {/* Sidebar */}
            <aside
                className={`bg-gray-800 text-white transition-all duration-300 ${collapsed ? 'w-16' : 'w-64'
                    }`}
            >
                {/* Sidebar header */}
                <div className="flex items-center justify-between h-16 px-4 border-b border-gray-700">
                    {!collapsed && <h1 className="text-xl font-semibold">Task Admin</h1>}
                    <button
                        onClick={() => setCollapsed(!collapsed)}
                        className="p-1 rounded-md hover:bg-gray-700"
                    >
                        {collapsed ? <Menu size={20} /> : <X size={20} />}
                    </button>
                </div>

                {/* Navigation */}
                <nav className="mt-5">
                    <ul>
                        {navItems.map((item) => (
                            <li key={item.path} className="mb-2">
                                <Link
                                    to={item.path}
                                    className={`flex items-center px-4 py-3 text-sm ${isActive(item.path)
                                        ? 'bg-blue-600 font-medium'
                                        : 'hover:bg-gray-700'
                                        }`}
                                >
                                    <span className="mr-3">{item.icon}</span>
                                    {!collapsed && <span>{item.name}</span>}
                                </Link>
                            </li>
                        ))}
                    </ul>
                </nav>
            </aside>

            {/* Main content */}
            <div className="flex flex-col flex-1 overflow-hidden">
                {/* Header */}
                <header className="bg-white shadow-sm z-10">
                    <div className="flex items-center justify-between h-16 px-6">
                        <div className="flex items-center w-64">
                            <div className="relative w-full">
                                <span className="absolute inset-y-0 left-0 flex items-center pl-2">
                                    <Search size={18} className="text-gray-400" />
                                </span>
                                <input
                                    type="text"
                                    placeholder="Search..."
                                    className="w-full py-2 pl-10 pr-4 text-sm bg-gray-100 border border-transparent rounded-md focus:outline-none focus:bg-white focus:border-gray-300"
                                />
                            </div>
                        </div>
                        <div className="flex items-center">
                            {/* Notifications */}
                            <button className="p-2 mr-4 text-gray-400 hover:text-gray-500">
                                <Bell size={20} />
                            </button>
                            {/* User profile */}
                            <div className="relative">
                                <button className="flex items-center text-sm text-gray-700 focus:outline-none">
                                    <img
                                        className="w-8 h-8 rounded-full"
                                        src="/api/placeholder/32/32"
                                        alt="User avatar"
                                    />
                                    {!collapsed && <span className="ml-2">Admin User</span>}
                                </button>
                            </div>
                        </div>
                    </div>
                </header>

                {/* Page content */}
                <main className="flex-1 overflow-y-auto p-6 bg-gray-100">
                    <div className="container mx-auto">
                        {/* Breadcrumbs */}
                        <div className="mb-6">
                            <h1 className="text-2xl font-semibold text-gray-800">Dashboard</h1>
                            <p className="text-sm text-gray-500">
                                Home / {location.pathname.split("/").filter(Boolean).pop()}
                            </p>
                        </div>

                        {/* Content outlet */}
                        <div className="bg-white rounded-lg shadow p-6">
                            <Outlet />
                        </div>
                    </div>
                </main>
            </div>
        </div>
    );
}