import type { Ingredient } from "@/api/types/getIngredientListResponse";
import { Button } from "@/components/ui/button";
import { Card, Table } from "@chakra-ui/react";
import { MdDelete, MdEdit } from "react-icons/md";

type IngredientListTableViewProps = {
	ingredients: Ingredient[];
};

export const IngredientListTableView = ({
	ingredients,
}: IngredientListTableViewProps) => {
	return (
		<Card.Root variant={"elevated"} minHeight={"500px"} overflowY={"auto"}>
			<Card.Header>
				<Card.Title fontSize={"xl"} fontWeight={"700"} color={"green.700"}>
					食材リスト
				</Card.Title>
			</Card.Header>
			<Card.Body pt={2}>
				<Table.Root>
					<Table.Header>
						<Table.Row>
							<Table.ColumnHeader fontWeight={"600"} color={"green.600"}>
								食材名
							</Table.ColumnHeader>
							<Table.ColumnHeader fontWeight={"600"} color={"green.600"}>
								種別
							</Table.ColumnHeader>
							<Table.ColumnHeader fontWeight={"600"} color={"green.600"}>
								操作
							</Table.ColumnHeader>
						</Table.Row>
					</Table.Header>
					<Table.Body>
						{ingredients.map((ingredient) => (
							<Table.Row key={ingredient.id}>
								<Table.Cell>{ingredient.name}</Table.Cell>
								{/* <Table.Cell>{ingredient.category}</Table.Cell> */}
								<Table.Cell display={"flex"} gap={2}>
									<Button
										color={"green.400"}
										borderColor={"green.400"}
										variant={"outline"}
									>
										<MdEdit />
									</Button>
									<Button
										color={"red.400"}
										borderColor={"red.400"}
										variant={"outline"}
									>
										<MdDelete />
									</Button>
								</Table.Cell>
							</Table.Row>
						))}
					</Table.Body>
				</Table.Root>
			</Card.Body>
		</Card.Root>
	);
};
