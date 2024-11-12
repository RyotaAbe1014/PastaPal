import { ApiClient } from "@/api/client";
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

type generateTokenAndGetUserRequest = {
	code: string;
};

type generateTokenAndGetUserResponse = {
	userId: string;
	avatarUrl: string;
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
	const api = ApiClient();
	const { code } = params;
	useEffect(() => {
		let ignore = false;
		const generateTokenAndGetUser = async () => {
			const user = await api.Post<
				generateTokenAndGetUserRequest,
				generateTokenAndGetUserResponse
			>("/auth/github/token", {
				code,
			});
			// TODO: ユーザー情報を保存する
			console.log(user);
			// TODO: tanstack routerのnavigateがあるはず
			// window.location.href = "/";
		};
		if (!ignore) {
			generateTokenAndGetUser();
		}

		return () => {
			ignore = true;
		};
	}, [code, api]);

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
