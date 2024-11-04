import type { Meta, StoryObj } from "@storybook/react";
import { IngredientsLayout } from ".";

const meta: Meta<typeof IngredientsLayout> = {
	component: IngredientsLayout,
	title: "Components/Layout/IngredientsLayout",
};

export default meta;
type Story = StoryObj<typeof IngredientsLayout>;

export const Default: Story = {
	args: {
		children: "Ingredients Layout",
	},
};
