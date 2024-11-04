import { Box } from "@chakra-ui/react";
import { Header } from "../Header";

export type BaseLayoutProps = {
	children: React.ReactNode;
};

export const BaseLayout = ({ children }: BaseLayoutProps) => {
	return (
		<Box minH={"100vh"} display={"flex"} flexDirection={"column"}>
			<Header path="ingredients" />
			<Box as="main" flex="1" p="4" bg={"green.100"}>
				{children}
			</Box>
		</Box>
	);
};
