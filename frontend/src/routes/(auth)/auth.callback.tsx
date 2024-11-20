import { ApiClient } from "@/api/client";
import { GenerateTokenAndGetUserRequest } from "@/api/types/generateTokenAndGetUserRequest";
import { GenerateTokenAndGetUserResponse } from "@/api/types/generateTokenAndGetUserResponse";
import {
	ProgressCircleRing,
	ProgressCircleRoot,
} from "@/components/ui/progress-circle";
import { Box, Heading } from "@chakra-ui/react";
import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { useCallback, useEffect, useRef } from "react";

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
	const navigate = useNavigate();
	const api = ApiClient();
	const { code } = params;

	const hasFetched = useRef(false); // 初回実行を記録するフラグ(StrictMode対策)

	const generateTokenAndGetUser = useCallback(async () => {
		const user = await api.Post<
			GenerateTokenAndGetUserRequest,
			GenerateTokenAndGetUserResponse
		>("/auth/github/token", { code });
		localStorage.setItem("userId", JSON.stringify(user.userId));
		navigate({ to: "/ingredients" });
	}, [api, code, navigate]);

	useEffect(() => {
		if (!hasFetched.current) {
			generateTokenAndGetUser();
			hasFetched.current = true; // 初回実行を記録
		}
	}, [generateTokenAndGetUser]);

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
