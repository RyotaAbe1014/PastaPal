import { Button } from "@/components/ui/button";
import { Field } from "@/components/ui/field";
import { Card, Input } from "@chakra-ui/react";

export const LoginForm = () => {
	return (
		<Card.Root variant={"elevated"} width={"500px"}>
			<Card.Header>
				<Card.Title fontSize={"xl"} fontWeight={"700"} color={"green.700"}>
					管理者ログイン
				</Card.Title>
			</Card.Header>
			<Card.Body display={"flex"} flexDirection={"column"} gap={8} pt={6}>
				<Field label="メールアドレス">
					<Input
						type={"text"}
						placeholder="example@example.com"
						backgroundColor={"green.100"}
					/>
				</Field>
				<Field label="パスワード">
					<Input
						type="password"
						placeholder="example@example.com"
						backgroundColor={"green.100"}
					/>
				</Field>
				<Button type="button" backgroundColor={"green.600"}>
					ログイン
				</Button>
			</Card.Body>
		</Card.Root>
	);
};
