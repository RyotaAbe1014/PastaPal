import {
	useGetIngredientCategoryList,
	useGetIngredientList,
} from "../../hooks";
import { IngredientListTableView } from "./view";

export const IngredientListTable = () => {
	const { ingredients, isLoading: isIngredientsLoading } =
		useGetIngredientList();
	const { ingredientCategories, isLoading: isIngredientCategoriesLoading } =
		useGetIngredientCategoryList();

	if (
		isIngredientsLoading ||
		isIngredientCategoriesLoading ||
		!ingredients ||
		!ingredientCategories
	) {
		return <div>Loading...</div>;
	}

	const onDeleteButtonClick = (ingredientId: string) => {
		console.log(ingredientId);
	};

	const onEditButtonClick = (ingredientId: string) => {
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
