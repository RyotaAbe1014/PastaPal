import type { Meta, StoryObj } from "@storybook/react";
import { IngredientCategoryCreateFormView } from ".";

const meta: Meta<typeof IngredientCategoryCreateFormView> = {
	component: IngredientCategoryCreateFormView,
	title: "Components/Ingredient/IngredientCategoryCreateForm",
};

export default meta;
type Story = StoryObj<typeof IngredientCategoryCreateFormView>;

export const Default: Story = {
	args: {},
};
