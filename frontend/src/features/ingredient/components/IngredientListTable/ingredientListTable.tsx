import { ApiClient, isApiError } from "@/api/client";
import type { NoContentResponse } from "@/api/types/common/noContentResponse";
import type { Ingredient } from "@/api/types/getIngredientListResponse";
import { useState } from "react";
import {
	useGetIngredientCategoryList,
	useGetIngredientList,
} from "../../hooks";
import { IngredientEditFormDialog } from "../IngredientEditFormDialog";
import { IngredientListTableView } from "./view";

export const IngredientListTable = () => {
	const api = ApiClient();
	const [isEditDialogOpen, setIsEditDialogOpen] = useState(false);
	const [targetIngredient, setTargetIngredient] = useState<Ingredient | null>(
		null,
	);

	const {
		ingredients,
		isLoading: isIngredientsLoading,
		mutate,
		isValidating: isIngredientsValidating,
	} = useGetIngredientList();
	const { ingredientCategories, isLoading: isIngredientCategoriesLoading } =
		useGetIngredientCategoryList();

	if (
		isIngredientsLoading ||
		isIngredientCategoriesLoading ||
		isIngredientsValidating ||
		!ingredients ||
		!ingredientCategories
	) {
		return <div>Loading...</div>;
	}

	const onDeleteButtonClick = async (ingredientId: string) => {
		const response = await api.Delete<undefined, NoContentResponse>(
			`/ingredients/${ingredientId}`,
		);
		if (isApiError(response)) {
			console.error(response);
			// TODO: エラー処理
			return;
		}
		mutate();
	};

	const onEditButtonClick = (ingredient: Ingredient) => {
		setTargetIngredient(ingredient);
		setIsEditDialogOpen(true);
	};

	return (
		<>
			<IngredientListTableView
				ingredients={ingredients}
				ingredientCategories={ingredientCategories}
				onDeleteButtonClick={onDeleteButtonClick}
				onEditButtonClick={onEditButtonClick}
			/>
			<IngredientEditFormDialog
				ingredient={targetIngredient}
				open={isEditDialogOpen}
				setOpen={setIsEditDialogOpen}
			/>
		</>
	);
};
