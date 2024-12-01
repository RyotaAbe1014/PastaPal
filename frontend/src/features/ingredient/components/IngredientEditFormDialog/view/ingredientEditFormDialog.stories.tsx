import type { Meta, StoryObj } from "@storybook/react";
import { IngredientEditFormDialogView } from ".";

const meta: Meta<typeof IngredientEditFormDialogView> = {
	component: IngredientEditFormDialogView,
	title: "Components/Ingredient/IngredientEditFormDialog",
};

export default meta;
type Story = StoryObj<typeof IngredientEditFormDialogView>;

export const Default: Story = {
	args: {
		open: true,
		setOpen: () => {},
		onConfirm: () => {},
		ingredientCategoryList: [
			{
				id: "1",
				name: "野菜",
			},
			{
				id: "2",
				name: "果物",
			},
			{
				id: "3",
				name: "肉",
			}
		],
		name: "トマト",
		setName: () => {},
		ingredientCategoryId: "1",
		setIngredientCategoryId: () => {},
	},
};
