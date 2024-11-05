import type { Meta, StoryObj } from "@storybook/react";
import { IngredientCreateForm } from ".";

const meta: Meta<typeof IngredientCreateForm> = {
	component: IngredientCreateForm,
	title: "Components/Layout/IngredientCreateForm",
};

export default meta;
type Story = StoryObj<typeof IngredientCreateForm>;

export const Default: Story = {
	args: {},
};
