import type { Meta, StoryObj } from "@storybook/react";
import { AccountButton } from ".";

const meta: Meta<typeof AccountButton> = {
	component: AccountButton,
	title: "Components/Layout/AccountButton",
};

export default meta;
type Story = StoryObj<typeof AccountButton>;

export const Dashboard: Story = {
	args: {},
};
