import { Box } from "@chakra-ui/react";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/(auth)/login")({
	component: () => <Login />,
});

const Login = () => {
	return (
		// 中央寄せ
		<Box
			as="main"
			height={"100vh"}
			display="flex"
			justifyContent="center"
			alignItems="center"
		>
			{/* TODO: ログインフォーム */}
			loginformです
		</Box>
	);
};
