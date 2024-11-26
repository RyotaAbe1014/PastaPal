import { ApiClient, isApiError } from "@/api/client";
import type { GetGitHubUrlResponse } from "@/api/types/getGitHubUrlResponse";
import { Button } from "@/components/ui/button";
import { Card } from "@chakra-ui/react";
import { FaGithub } from "react-icons/fa";

export const LoginForm = () => {
	const api = ApiClient();
	const handleClick = async () => {
		const response = await api.Get<undefined, GetGitHubUrlResponse>(
			"/auth/github/url",
		);

		if (isApiError(response)) {
			// TODO: ダイアログ側にエラーを表示する
			return
		}

		window.location.href = response.data.url;
	};
	return (
		<Card.Root variant={"elevated"} width={"500px"}>
			<Card.Header>
				<Card.Title fontSize={"xl"} fontWeight={"700"} color={"green.700"}>
					管理者ログイン
				</Card.Title>
			</Card.Header>
			<Card.Body display={"flex"} flexDirection={"column"} gap={8} pt={6}>
				<Button type="button" backgroundColor={"black"} onClick={handleClick}>
					<FaGithub />
					GitHubでログイン
				</Button>
			</Card.Body>
		</Card.Root>
	);
};
