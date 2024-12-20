import { Box, Heading } from "@chakra-ui/react";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_app/recipes")({
	component: Recipes,
});

function Recipes() {
	return (
		<Box as="div" display={"flex"} flexDirection={"column"} gap={4}>
			<Heading as="h1" size="lg" color={"green.600"}>
				パスタレシピ一覧
			</Heading>
			{/* TODO: レシピ管理 */}
		</Box>
	);
}
