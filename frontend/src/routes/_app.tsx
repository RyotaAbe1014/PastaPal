import { ApiClient } from "@/api/client";
import type { AuthStatusResponse } from "@/api/types/authStatusResponse";
import { BaseLayout } from "@/components/layout/BaseLayout";
import { Outlet, createFileRoute, redirect } from "@tanstack/react-router";

export const Route = createFileRoute("/_app")({
	component: AppLayout,
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
	},
});

function AppLayout() {
	return (
		<BaseLayout>
			<Outlet />
		</BaseLayout>
	);
}
