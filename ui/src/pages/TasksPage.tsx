import graphql from 'babel-plugin-relay/macro';
import { useLazyLoadQuery } from 'react-relay';
import WebTable, { Column } from '../components/WebTable';
import { TasksPageQuery } from './__generated__/TasksPageQuery.graphql';

// Define Task interface
interface Task {
    id: string;
    name: string;
    status: 'Pending' | 'In Progress' | 'Completed' | 'Failed';
    priority: 'Low' | 'Medium' | 'High' | 'Critical';
    createdAt: string;
    dueDate: string;
}

const query = graphql`
    query TasksPageQuery {
        entTasks {
            edges {
                node {
                    id
                    handler
                    parameter
                }
            }
        }
    }
`;

export default function TasksPage() {
    const data = useLazyLoadQuery<TasksPageQuery>(query, {});
    const tasks: Task[] = (data.entTasks.edges ?? []).map((edge): Task => {
        return {
            id: edge?.node?.id ?? '',
            name: edge?.node?.handler ?? '',
            status: 'In Progress',
            priority: 'High',
            createdAt: '2025-03-28T09:15:00Z',
            dueDate: '2025-04-12T18:00:00Z',
        }
    })

    // Column definitions for tasks with TypeScript typing
    const taskColumns: Column<Task>[] = [
        { field: 'id', headerName: 'Task ID', width: 120 },
        { field: 'name', headerName: 'Task Name', width: 200 },
        {
            field: 'status',
            headerName: 'Status',
            width: 120,
            renderCell: (row: Task) => {
                const statusColors: Record<Task['status'], string> = {
                    'Completed': 'bg-green-100 text-green-800',
                    'In Progress': 'bg-blue-100 text-blue-800',
                    'Pending': 'bg-yellow-100 text-yellow-800',
                    'Failed': 'bg-red-100 text-red-800'
                };

                return (
                    <span className={`px-2 py-1 rounded-full text-xs font-medium ${statusColors[row.status]}`}>
                        {row.status}
                    </span>
                );
            }
        },
        {
            field: 'priority',
            headerName: 'Priority',
            width: 100,
            renderCell: (row: Task) => {
                const priorityColors: Record<Task['priority'], string> = {
                    'Low': 'bg-gray-100 text-gray-800',
                    'Medium': 'bg-blue-100 text-blue-800',
                    'High': 'bg-orange-100 text-orange-800',
                    'Critical': 'bg-red-100 text-red-800'
                };

                return (
                    <span className={`px-2 py-1 rounded-full text-xs font-medium ${priorityColors[row.priority]}`}>
                        {row.priority}
                    </span>
                );
            }
        },
        { field: 'assignedTo', headerName: 'Assigned To', width: 150 },
        {
            field: 'createdAt',
            headerName: 'Created',
            width: 150,
            renderCell: (row: Task) => {
                const date = new Date(row.createdAt);
                return date.toLocaleDateString();
            }
        },
        {
            field: 'dueDate',
            headerName: 'Due Date',
            width: 150,
            renderCell: (row: Task) => {
                const date = new Date(row.dueDate);
                return date.toLocaleDateString();
            }
        },
        {
            field: 'actions',
            headerName: 'Actions',
            width: 100,
            renderCell: (row: Task) => (
                <div className="flex space-x-2">
                    <button
                        className="p-1 text-blue-600 hover:text-blue-800"
                        onClick={() => console.log(`Edit task ${row.id}`)}
                    >
                        Edit
                    </button>
                    <button
                        className="p-1 text-red-600 hover:text-red-800"
                        onClick={() => console.log(`Delete task ${row.id}`)}
                    >
                        Delete
                    </button>
                </div>
            )
        }
    ];

    return (
        <div className="space-y-8">
            <WebTable
                rows={tasks}
                columns={taskColumns}
                label="Task Management"
                pageSize={5}
            />
        </div>
    );
};
