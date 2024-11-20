import { ApiClient } from "@/api/client";
import { createFileRoute, redirect } from "@tanstack/react-router";

type AuthStatusResponse = {
	isAuthenticated: boolean;
};

export const Route = createFileRoute("/")({
	beforeLoad: async () => {
		const api = ApiClient();
		const response = await api.Get<undefined, AuthStatusResponse>(
			"/auth/github/status",
		);
		if (!response.isAuthenticated) {
			// APIはcookieのトークンを見ているので、ローカルストレージのユーザー名だけでログイン済みとみなす
			throw redirect({
				to: "/login",
			});
		}
		throw redirect({
			to: "/ingredients",
		});
	},
});
