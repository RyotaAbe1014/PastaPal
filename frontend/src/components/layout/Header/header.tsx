import { Button } from "@/components/ui/button";
import { Box, Heading } from "@chakra-ui/react";

export type HeaderProps = {
	path: string;
};

export const Header = ({ path }: HeaderProps) => {
	return (
		<Box
			as="header"
			bg="green.600"
			color="white"
			p={4}
			display={"flex"}
			justifyContent={"space-between"}
			alignItems={"center"}
		>
			<Heading as="h1" size="lg">
				PastaPal
			</Heading>
			<Box
				display={"flex"}
				alignItems={"center"}
				justifyContent={"space-around"}
			>
				<Box mr={4}>
					{/* ダッシュボード遷移ボタン */}
					<Button
						bg={getButtonBackGroundColor(path, "dashboard")}
						color={getButtonTextColor(path, "dashboard")}
					>
						ダッシュボード
					</Button>
				</Box>
				<Box mr={4}>
					{/* レシピ遷移ボタン */}
					<Button
						bg={getButtonBackGroundColor(path, "recipe")}
						color={getButtonTextColor(path, "recipe")}
					>
						レシピ
					</Button>
				</Box>
				<Box mr={4}>
					{/* 食材遷移ボタン */}
					<Button
						bg={getButtonBackGroundColor(path, "ingredients")}
						color={getButtonTextColor(path, "ingredients")}
					>
						食材
					</Button>
				</Box>
				<Box mr={4}>{/* アカウントアイコン */}</Box>
			</Box>
		</Box>
	);
};

const getButtonBackGroundColor = (path: string, targetPath: string) => {
	return path === targetPath ? "white" : "green.600";
};

const getButtonTextColor = (path: string, targetPath: string) => {
	return path === targetPath ? "green.600" : "white";
};