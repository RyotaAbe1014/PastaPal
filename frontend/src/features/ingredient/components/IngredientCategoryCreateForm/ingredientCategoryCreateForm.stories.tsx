import type { Meta, StoryObj } from "@storybook/react";
import { IngredientCategoryCreateForm } from ".";

const meta: Meta<typeof IngredientCategoryCreateForm> = {
	component: IngredientCategoryCreateForm,
	title: "Components/Layout/IngredientCategoryCreateForm",
};

export default meta;
type Story = StoryObj<typeof IngredientCategoryCreateForm>;

export const Default: Story = {
	args: {},
};
