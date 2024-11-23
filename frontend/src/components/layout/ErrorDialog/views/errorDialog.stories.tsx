import { Button } from "@/components/ui/button";
import type { Meta, StoryObj } from "@storybook/react";
import { useState } from "react";
import { ErrorDialogView } from ".";

const meta: Meta<typeof ErrorDialogView> = {
	component: ErrorDialogView,
	title: "Components/Layout/ErrorDialog",
};

export default meta;
type Story = StoryObj<typeof ErrorDialogView>;

export const Default: Story = {
	args: {},
	render: () => <ErrorDialogWrapper />,
};

const ErrorDialogWrapper = () => {
	const [open, setOpen] = useState(false);
	const onConfirm = () => {
		alert("Confirmed");
		setOpen(false);
	};

	const onCancel = () => {
		setOpen(false);
	};

	return (
		<>
			<Button onClick={() => setOpen(true)}>Open Error Dialog</Button>
			<ErrorDialogView open={open} onConfirm={onConfirm} onCancel={onCancel} />
		</>
	);
};
