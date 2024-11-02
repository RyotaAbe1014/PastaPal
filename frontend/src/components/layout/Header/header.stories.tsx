import type { Meta, StoryObj } from "@storybook/react";
import { Header } from "./header";

const meta: Meta<typeof Header> = {
	component: Header,
	title: "Components/Layout/Header",
};

export default meta;
type Story = StoryObj<typeof Header>;

export const Dashboard: Story = {
	args: {
		path: "dashboard",
	},
};

export const Recipe: Story = {
	args: {
		path: "recipe",
	},
};

export const Ingredients: Story = {
	args: {
		path: "ingredients",
	},
};
