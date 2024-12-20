import type { IngredientCategory } from "@/api/types/getIngredientCategoryListResponse";
import { Button } from "@/components/ui/button";
import {
	DialogBody,
	DialogCloseTrigger,
	DialogContent,
	DialogFooter,
	DialogHeader,
	DialogRoot,
	DialogTitle,
} from "@/components/ui/dialog";
import {
	SelectContent,
	SelectItem,
	SelectRoot,
	SelectTrigger,
	SelectValueText,
} from "@/components/ui/select";
import { Input, createListCollection } from "@chakra-ui/react";
import { useMemo, useRef } from "react";

export type IngredientEditFormDialogViewProps = {
	open: boolean;
	setOpen: (open: boolean) => void;
	onConfirm: () => void;
	ingredientCategoryList: IngredientCategory[];
	name: string;
	setName: (name: string) => void;
	ingredientCategoryId: string;
	setIngredientCategoryId: (ingredientCategoryId: string) => void;
};

export const IngredientEditFormDialogView = ({
	open,
	setOpen,
	onConfirm,
	ingredientCategoryList,
	name,
	setName,
	ingredientCategoryId,
	setIngredientCategoryId,
}: IngredientEditFormDialogViewProps) => {
	// ダイアログ内でselectを使用するためには、refを使用する必要がある
	// https://www.chakra-ui.com/docs/components/select#within-dialog
	const contentRef = useRef<HTMLDivElement>(null);

	const items = useMemo(() => {
		return createListCollection({
			items: ingredientCategoryList.map((category) => ({
				label: category.name,
				value: category.id.toString(),
			})),
		});
	}, [ingredientCategoryList]);

	return (
		<DialogRoot
			lazyMount
			open={open}
			onOpenChange={(e) => setOpen(e.open)}
			placement={"center"}
		>
			<DialogContent>
				<DialogHeader>
					<DialogTitle>食材情報更新</DialogTitle>
				</DialogHeader>
				<DialogBody
					ref={contentRef}
					display={"flex"}
					flexDirection={"row"}
					gap={3}
				>
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
						<SelectContent portalRef={contentRef}>
							{items.items.map((item) => (
								<SelectItem item={item} key={item.value}>
									{item.label}
								</SelectItem>
							))}
						</SelectContent>
					</SelectRoot>
				</DialogBody>
				<DialogFooter>
					<Button variant="outline" onClick={() => setOpen(false)}>
						キャンセル
					</Button>
					<Button backgroundColor={"green.600"} onClick={onConfirm}>
						保存
					</Button>
				</DialogFooter>
				<DialogCloseTrigger />
			</DialogContent>
		</DialogRoot>
	);
};
