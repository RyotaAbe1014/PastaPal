import { Box } from "@chakra-ui/react";
import { Header } from "../Header";

export type BaseLayoutProps = {
	children: React.ReactNode;
	path: string;
};

export const BaseLayout = ({ children, path }: BaseLayoutProps) => {
	return (
		<Box minH={"100vh"} display={"flex"} flexDirection={"column"}>
			<Header path={path} />
			<Box as="main" flex="1" p="4" bg={"green.100"}>
				{children}
			</Box>
		</Box>
	);
};
