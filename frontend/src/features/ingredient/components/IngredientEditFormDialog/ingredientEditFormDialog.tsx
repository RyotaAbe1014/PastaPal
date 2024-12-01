import { ApiClient, isApiError } from "@/api/client";
import type { Ingredient } from "@/api/types/getIngredientListResponse";
import type { UpdateIngredientRequest } from "@/api/types/updateIngredientRequest";
import type { UpdateIngredientResponse } from "@/api/types/updateIngredientResponse";
import { useEffect, useState } from "react";
import {
	useGetIngredientCategoryList,
	useGetIngredientList,
} from "../../hooks";
import { IngredientEditFormDialogView } from "./view";

type IngredientEditFormDialogProps = {
	ingredient: Ingredient | null;
	open: boolean;
	setOpen: (open: boolean) => void;
};

export const IngredientEditFormDialog = ({
	ingredient,
	open,
	setOpen,
}: IngredientEditFormDialogProps) => {
	const api = ApiClient();
	const [name, setName] = useState<string>("");
	const [ingredientCategoryId, setIngredientCategoryId] = useState<string>("");

	const { ingredientCategories, isLoading } = useGetIngredientCategoryList();
	const { mutate } = useGetIngredientList();

	const onConfirm = async () => {
		if (!ingredientCategoryId || !name || !ingredient) {
			// TODO: エラーをダイアログで表示する
			return;
		}

		const response = await api.Put<
			UpdateIngredientRequest,
			UpdateIngredientResponse
		>(`/ingredients/${ingredient.id}`, {
			name,
			ingredientCategoryId: ingredientCategoryId,
		});

		if (isApiError(response)) {
			console.error(response.message);
			// TODO: エラーをダイアログで表示する
			return;
		}

		mutate();
		setOpen(false);
	};

	useEffect(() => {
		if (!ingredient) {
			return;
		}
		setName(ingredient.name);
		setIngredientCategoryId(ingredient.ingredientCategoryId);
	}, [ingredient]);

	if (isLoading || !ingredientCategories) {
		return <div>Loading...</div>;
	}

	return (
		<IngredientEditFormDialogView
			open={open}
			setOpen={setOpen}
			onConfirm={onConfirm}
			ingredientCategoryList={ingredientCategories}
			name={name}
			setName={setName}
			ingredientCategoryId={ingredientCategoryId}
			setIngredientCategoryId={setIngredientCategoryId}
		/>
	);
};
