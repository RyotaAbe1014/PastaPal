import type { Meta, StoryObj } from "@storybook/react";
import { IngredientCreateForm } from ".";

const meta: Meta<typeof IngredientCreateForm> = {
	component: IngredientCreateForm,
	title: "Components/Ingredient/IngredientCreateForm",
};

export default meta;
type Story = StoryObj<typeof IngredientCreateForm>;

export const Default: Story = {
	args: {},
};
