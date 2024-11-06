import { Box } from "@chakra-ui/react";
import { useLocation } from "@tanstack/react-router";
import { Header } from "../Header";

export type BaseLayoutProps = {
	children: React.ReactNode;
};

export const BaseLayout = ({ children }: BaseLayoutProps) => {
	const location = useLocation();
	return (
		<Box minH={"100vh"} display={"flex"} flexDirection={"column"}>
			<Header path={location.pathname} />
			<Box as="main" flex="1" p="4" bg={"green.100"}>
				{children}
			</Box>
		</Box>
	);
};
