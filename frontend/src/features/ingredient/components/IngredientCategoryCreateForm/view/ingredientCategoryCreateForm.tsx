import { Button } from "@/components/ui/button";
import { Card, Input } from "@chakra-ui/react";
import { BsPlus } from "react-icons/bs";

type IngredientCategoryCreateFormViewProps = {
	onConfirm: () => void;
	name: string;
	setName: (name: string) => void;
};

export const IngredientCategoryCreateFormView = ({
	onConfirm,
	name,
	setName,
}: IngredientCategoryCreateFormViewProps) => {
	return (
		<Card.Root variant={"elevated"} maxWidth={"900px"}>
			<Card.Header>
				<Card.Title fontSize={"xl"} fontWeight={"700"} color={"green.700"}>
					種別の追加
				</Card.Title>
			</Card.Header>
			<Card.Body display={"flex"} flexDirection={"row"} gap={"3"} pt={2}>
				<Input
					placeholder="新しい種別"
					backgroundColor={"green.100"}
					value={name}
					onChange={(e) => setName(e.target.value)}
				/>
				<Button backgroundColor={"green.600"} onClick={onConfirm}>
					<BsPlus />
					追加
				</Button>
			</Card.Body>
		</Card.Root>
	);
};
