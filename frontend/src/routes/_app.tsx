import { ApiClient } from "@/api/client";
import { BaseLayout } from "@/components/layout/BaseLayout";
import { Outlet, createFileRoute, redirect } from "@tanstack/react-router";

export const Route = createFileRoute("/_app")({
	component: AppLayout,
	beforeLoad: async ({ location }) => {
		const api = ApiClient();
		const isAuthenticated = await api.Get<boolean>("/auth/github/status");
		if (!isAuthenticated) {
			// APIはcookieのトークンを見ているので、ローカルストレージのユーザー名だけでログイン済みとみなす
			throw redirect({
				to: "/login",
				search: {
					redirect: location.href,
				},
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
