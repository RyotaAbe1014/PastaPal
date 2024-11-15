import type { Meta, StoryObj } from "@storybook/react";
import { IngredientListTable } from ".";

const meta: Meta<typeof IngredientListTable> = {
	component: IngredientListTable,
	title: "Components/Ingredient/IngredientListTable",
};

export default meta;
type Story = StoryObj<typeof IngredientListTable>;

export const Default: Story = {
	args: {},
};
