import { Button } from "@/components/ui/button";
import {
	DialogBody,
	DialogCloseTrigger,
	DialogContent,
	DialogFooter,
	DialogHeader,
	DialogRoot,
	DialogTitle,
} from "@/components/ui/dialog";

export type ErrorDialogProps = {
	open: boolean;
	title?: string;
	message?: string;
	onConfirm?: () => void;
	onCancel: () => void;
};

export const ErrorDialogView = ({
	open,
	title,
	message,
	onConfirm,
	onCancel,
}: ErrorDialogProps) => {
	return (
		<DialogRoot open={open}>
			<DialogContent>
				<DialogHeader>
					<DialogTitle>{title || "Error"}</DialogTitle>
				</DialogHeader>
				<DialogBody>
					<p>{message || "エラーが発生しました。もう一度お試しください。"}</p>
				</DialogBody>
				<DialogFooter>
					<Button variant="outline" onClick={onCancel}>
						Cancel
					</Button>
					<Button onClick={onConfirm}>OK</Button>
				</DialogFooter>
				<DialogCloseTrigger />
			</DialogContent>
		</DialogRoot>
	);
};
