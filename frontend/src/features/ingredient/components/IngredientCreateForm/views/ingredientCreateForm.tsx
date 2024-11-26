import type { IngredientCategory } from "@/api/types/getIngredientCategoryListResponse";
import { Button } from "@/components/ui/button";
import {
	SelectContent,
	SelectItem,
	SelectRoot,
	SelectTrigger,
	SelectValueText,
} from "@/components/ui/select";
import { Card, Input, createListCollection } from "@chakra-ui/react";
import { useMemo } from "react";
import { BsPlus } from "react-icons/bs";

export type IngredientCreateFormViewProps = {
	ingredientCategoryList: IngredientCategory[];
	name: string;
	setName: (name: string) => void;
	ingredientCategoryId: string;
	setIngredientCategoryId: (ingredientCategoryId: string) => void;
	onConfirm: () => void;
};

export const IngredientCreateFormView = ({
	ingredientCategoryList,
	name,
	setName,
	ingredientCategoryId,
	setIngredientCategoryId,
	onConfirm,
}: IngredientCreateFormViewProps) => {
	const items = useMemo(() => {
		return createListCollection({
			items: ingredientCategoryList.map((category) => ({
				label: category.name,
				value: category.id.toString(),
			})),
		});
	}, [ingredientCategoryList]);

	return (
		<Card.Root variant={"elevated"} maxWidth={"900px"}>
			<Card.Header>
				<Card.Title fontSize={"xl"} fontWeight={"700"} color={"green.700"}>
					食材の追加
				</Card.Title>
			</Card.Header>
			<Card.Body display={"flex"} flexDirection={"row"} gap={"3"} pt={2}>
				<Input
					placeholder="食材名"
					backgroundColor={"green.100"}
					flex="6"
					value={name}
					onChange={(e) => setName(e.target.value)}
				/>
				<SelectRoot
					size={"md"}
					collection={items}
					backgroundColor={"green.100"}
					flex={"3"}
					value={[ingredientCategoryId]}
					onValueChange={(e) => setIngredientCategoryId(e.value[0])}
				>
					<SelectTrigger>
						<SelectValueText placeholder="種別を選択" />
					</SelectTrigger>
					<SelectContent>
						{items.items.map((item) => (
							<SelectItem item={item} key={item.value}>
								{item.label}
							</SelectItem>
						))}
					</SelectContent>
				</SelectRoot>
				<Button backgroundColor={"green.600"} onClick={onConfirm}>
					<BsPlus />
					追加
				</Button>
			</Card.Body>
		</Card.Root>
	);
};
