import { ApiClient } from "@/api/client";
import { Button } from "@/components/ui/button";
import {
	PopoverArrow,
	PopoverBody,
	PopoverContent,
	PopoverRoot,
	PopoverTrigger,
} from "@/components/ui/popover";
import { IconButton, Stack, Text } from "@chakra-ui/react";
import { useNavigate } from "@tanstack/react-router";
import { useState } from "react";
import { VscAccount } from "react-icons/vsc";

export const AccountButton = () => {
	const navigate = useNavigate();
	const api = ApiClient();
	const [open, setOpen] = useState(false);
	const userId = localStorage.getItem("userId");
	const handleLogout = async () => {
		localStorage.removeItem("userId");
		await api.Post("/auth/github/logout");
		navigate({ to: "/login" });
	};

	return (
		<PopoverRoot open={open} onOpenChange={(e) => setOpen(e.open)}>
			<PopoverTrigger asChild>
				<IconButton rounded="full" aria-label="アカウント" bg="green.600">
					<VscAccount />
				</IconButton>
			</PopoverTrigger>
			<PopoverContent>
				<PopoverArrow />
				<PopoverBody>
					<Stack gap="4">
						<Text>ユーザー名: {userId}</Text>
						<Button type="button" bg="green.600" onClick={handleLogout}>
							ログアウト
						</Button>
					</Stack>
				</PopoverBody>
			</PopoverContent>
		</PopoverRoot>
	);
};
