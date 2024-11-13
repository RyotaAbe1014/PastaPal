import { IconButton } from "@chakra-ui/react";
import { VscAccount } from "react-icons/vsc";

export const AccountButton = () => {
  return (
    <IconButton rounded="full" aria-label="アカウント" bg="green.600">
      <VscAccount />
    </IconButton>
  );
}
