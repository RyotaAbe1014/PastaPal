import {
	ProgressCircleRing,
	ProgressCircleRoot,
} from "@/components/ui/progress-circle";
import { Box, Heading } from "@chakra-ui/react";
import { createFileRoute } from "@tanstack/react-router";
import { useEffect } from "react";

type AuthCallbackProps = {
	code: string;
	state: string;
};

export const Route = createFileRoute("/(auth)/auth/callback")({
	validateSearch: (search: Record<string, unknown>): AuthCallbackProps => {
		const code = search?.code;
		const state = search?.state;
		if (typeof code !== "string" || typeof state !== "string") {
			throw new Error("Invalid params");
		}
		return { code, state };
	},
	component: () => <AuthCallback />,
});

const AuthCallback = () => {
	const params = Route.useSearch();
	const { code, state } = params;

	useEffectF(() => {
		// const generate
	}, [code, state]);

	return (
		<Box
			as="main"
			height={"100vh"}
			display="flex"
			justifyContent="center"
			flexDirection={"column"}
			gap={6}
			alignItems="center"
			bgColor={"green.100"}
		>
			<ProgressCircleRoot value={null} size="xl" colorPalette={"green"}>
				<ProgressCircleRing cap="round" />
			</ProgressCircleRoot>
			<Heading as="h1" size="lg" color={"green.700"}>
				ユーザー情報を取得しています...
			</Heading>
		</Box>
	);
};
