import { ApiClient } from "@/api/client";

const api = ApiClient();

export const fetcher = async <RequestType = undefined, ResponseType = unknown>(path: string, params?: RequestType) => {
  const response = await api.Get<RequestType, ResponseType>(path, params);

  if (response.type === 'error') {
    // SWRのエラーとして扱うため、エラーを投げる
    // カスタムエラーをスロー
    const error = new Error(response.message) as Error & { status: string; title?: string };
    error.status = response.status;
    error.title = response.title;
    throw error;
  }

  return response.data;
};
