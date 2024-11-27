import type { Meta, StoryObj } from "@storybook/react";
import { IngredientListTableView } from ".";

const meta: Meta<typeof IngredientListTableView> = {
	component: IngredientListTableView,
	title: "Components/Ingredient/IngredientListTable",
};

export default meta;
type Story = StoryObj<typeof IngredientListTableView>;

export const Default: Story = {
	args: {},
};
