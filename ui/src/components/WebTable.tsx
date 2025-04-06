import React, { useState } from "react";
import { ChevronDown, ChevronUp, Search, Filter, Download } from "lucide-react";

export interface Column<T> {
  field: keyof T | string;
  headerName: string;
  width?: number;
  renderCell?: (row: T) => React.ReactNode;
}

export interface Props<T> {
  rows: T[];
  columns: Column<T>[];
  label: string;
  isLoading?: boolean;
  pageSize?: number;
  emptyMessage?: string;
}

function WebTable<T>({
  rows = [],
  columns = [],
  label,
  isLoading = false,
  pageSize = 10,
  emptyMessage = "No data available",
}: Props<T>): React.ReactElement {
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [sortField, setSortField] = useState<string | null>(null);
  const [sortDirection, setSortDirection] = useState<"asc" | "desc">("asc");
  const [searchTerm, setSearchTerm] = useState<string>("");

  return (
    <div className="bg-white rounded-lg shadow border border-gray-200">
      {/* Table header with title and controls */}
      <div className="px-6 py-4 border-b border-gray-200 flex items-center justify-between flex-wrap gap-4">
        <h2 className="text-lg font-medium text-gray-800">{label}</h2>
        <div className="flex items-center space-x-4">
          {/* Search */}
          <div className="relative">
            <span className="absolute inset-y-0 left-0 flex items-center pl-2">
              <Search size={16} className="text-gray-400" />
            </span>
            <input
              type="text"
              placeholder="Search..."
              value={searchTerm}
              onChange={(e) => {
                setSearchTerm(e.target.value);
                setCurrentPage(1); // Reset to first page when searching
              }}
              className="py-2 pl-10 pr-4 text-sm bg-gray-100 border border-transparent rounded-md w-64 focus:outline-none focus:bg-white focus:border-gray-300"
            />
          </div>

          {/* Actions */}
          <button className="p-2 text-gray-500 hover:text-gray-700 bg-gray-100 rounded">
            <Filter size={16} />
          </button>
          <button className="p-2 text-gray-500 hover:text-gray-700 bg-gray-100 rounded">
            <Download size={16} />
          </button>
        </div>
      </div>

      {/* Table content */}
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              {columns.map((column) => (
                <th
                  key={String(column.field)}
                  className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider cursor-pointer hover:bg-gray-100"
                  onClick={() => { }}
                >
                  <div className="flex items-center">
                    {column.headerName}
                    {sortField === column.field && (
                      <span className="ml-1">
                        {sortDirection === "asc" ? (
                          <ChevronUp size={14} />
                        ) : (
                          <ChevronDown size={14} />
                        )}
                      </span>
                    )}
                  </div>
                </th>
              ))}
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {isLoading ? (
              <tr>
                <td
                  colSpan={columns.length}
                  className="px-6 py-4 text-center text-sm text-gray-500"
                >
                  Loading...
                </td>
              </tr>
            ) : rows.length === 0 ? (
              <tr>
                <td
                  colSpan={columns.length}
                  className="px-6 py-4 text-center text-sm text-gray-500"
                >
                  {emptyMessage}
                </td>
              </tr>
            ) : (
              rows.map((row, rowIndex) => (
                <tr key={rowIndex} className="hover:bg-gray-50">
                  {columns.map((column) => (
                    <td
                      key={String(column.field)}
                      className="px-6 py-4 whitespace-nowrap text-sm text-gray-500"
                    >
                      {column.renderCell
                        ? column.renderCell(row)
                        : (row[column.field as keyof T] as React.ReactNode)}
                    </td>
                  ))}
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default WebTable;
