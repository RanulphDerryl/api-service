interface ApiResponse<T> {
    data: T;
    status: number;
    statusText: string;
    headers: Record<string, string>;
}

interface ApiError {
    message: string;
    status?: number;
    code?: string;
}

interface RequestOptions {
    method?: 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH';
    headers?: Record<string, string>;
    body?: any;
    params?: Record<string, string | number | boolean>;
}

interface PaginatedResponse<T> {
    items: T[];
    total: number;
    page: number;
    pageSize: number;
}

interface User {
    id: string;
    name: string;
    email: string;
    createdAt: string;
    updatedAt: string;
}

interface AuthResponse {
    token: string;
    user: User;
}

type ApiFunction<T> = (options?: RequestOptions) => Promise<ApiResponse<T>>;

export type {
    ApiResponse,
    ApiError,
    RequestOptions,
    PaginatedResponse,
    User,
    AuthResponse,
    ApiFunction
};