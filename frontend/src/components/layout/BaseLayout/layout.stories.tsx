import type { Meta, StoryObj } from "@storybook/react";
import { BaseLayout } from ".";

const meta: Meta<typeof BaseLayout> = {
	component: BaseLayout,
	title: "Components/Layout/BaseLayout",
};

export default meta;
type Story = StoryObj<typeof BaseLayout>;

export const Default: Story = {
	args: {
		children: "Ingredients Layout",
		path: "ingredients",
	},
};
