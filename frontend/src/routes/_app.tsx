import { BaseLayout } from "@/components/layout/BaseLayout";
import { Outlet, createFileRoute, redirect } from "@tanstack/react-router";

export const Route = createFileRoute("/_app")({
	component: AppLayout,
	beforeLoad: async ({ location }) => {
		const userId = localStorage.getItem("userId");
		if (!userId) {
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
