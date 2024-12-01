import type { Ingredient } from "@/api/types/getIngredientListResponse";
import type { Meta, StoryObj } from "@storybook/react";
import { IngredientListTableView } from ".";

const meta: Meta<typeof IngredientListTableView> = {
	component: IngredientListTableView,
	title: "Components/Ingredient/IngredientListTable",
};

export default meta;
type Story = StoryObj<typeof IngredientListTableView>;

export const Default: Story = {
	args: {
		onDeleteButtonClick: (ingredientId: string) => {
			console.log(ingredientId);
		},

		onEditButtonClick: (ingredient: Ingredient) => {
			console.log(ingredient);
		},
		ingredients: [
			{
				id: "1",
				name: "トマト",
				ingredientCategoryId: "1",
			},
			{
				id: "2",
				name: "キュウリ",
				ingredientCategoryId: "1",
			},
		],
		ingredientCategories: [
			{
				id: "1",
				name: "野菜",
			},
		],
	},
};
