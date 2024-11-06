import { Button } from "@/components/ui/button";
import { Table } from "@chakra-ui/react";
import { MdDelete, MdEdit } from "react-icons/md";

export const IngredientListTable = () => {
	// TODO: ここはAPIから取得するようにする
	const data = [
		{
			id: 1,
			name: "トマト",
			category: "野菜",
		},
		{
			id: 2,
			name: "玉ねぎ",
			category: "野菜",
		},
		{
			id: 3,
			name: "鶏肉",
			category: "肉",
		},
	];
	return (
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
				{data.map((item) => (
					<Table.Row key={item.id}>
						<Table.Cell>{item.name}</Table.Cell>
						<Table.Cell>{item.category}</Table.Cell>
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
	);
};
