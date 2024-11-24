import type { Meta, StoryObj } from "@storybook/react";
import { IngredientCreateFormView } from ".";

const meta: Meta<typeof IngredientCreateFormView> = {
	component: IngredientCreateFormView,
	title: "Components/Ingredient/IngredientCreateForm",
};

export default meta;
type Story = StoryObj<typeof IngredientCreateFormView>;

export const Default: Story = {
	args: {
		ingredientCategoryList: [],
	},
};
