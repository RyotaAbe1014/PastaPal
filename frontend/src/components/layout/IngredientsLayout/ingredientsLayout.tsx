import { Box } from "@chakra-ui/react";
import { Header } from "../Header";


export type IngredientsLayoutProps = {
  children: React.ReactNode;
};

export const IngredientsLayout = ({children}: IngredientsLayoutProps) => {
  return (
    <Box minH={"100vh"} display={"flex"} flexDirection={"column"}>
      <Header path="ingredients" />
      <Box as="main" flex="1" p="4" bg={"green.100"}>
        {children}
      </Box>
    </Box>
  )
};