import { BaseLayout } from "@/components/layout/BaseLayout";
import { Outlet, createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/(app)/_recipes")({
	component: Recipes,
});

function Recipes() {
	return (
		<BaseLayout path={"recipes"}>
			<Outlet />
		</BaseLayout>
	);
}
