import { ApiClient, isApiError } from "@/api/client";
import type { NoContentResponse } from "@/api/types/common/noContentResponse";
import {
	useGetIngredientCategoryList,
	useGetIngredientList,
} from "../../hooks";
import { IngredientListTableView } from "./view";

export const IngredientListTable = () => {
	const api = ApiClient();
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

	const onEditButtonClick = (ingredientId: string) => {
		// TODO: 編集画面に遷移
		console.log(ingredientId);
	};

	return (
		<IngredientListTableView
			ingredients={ingredients}
			ingredientCategories={ingredientCategories}
			onDeleteButtonClick={onDeleteButtonClick}
			onEditButtonClick={onEditButtonClick}
		/>
	);
};
