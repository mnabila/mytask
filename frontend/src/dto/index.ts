export interface AuthRequest {
  email: string;
  password: string;
}

export interface UserResponse {
  id: string;
  name: string;
  email: string;
}

export interface TodoResponse {
  id: number;
  createdAt: string;
  updatedAt: string;
  task: string;
  description: string;
}

export interface TodoRquest {
  task: string;
}
