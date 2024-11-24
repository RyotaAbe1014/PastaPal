import { useIngredientCreateForm } from "./hooks";
import { IngredientCreateFormView } from "./views";

export const IngredientCreateForm = () => {
	const { ingredientCategoryList, isLoading } = useIngredientCreateForm();

	if (isLoading || typeof ingredientCategoryList === "undefined") {
		return <div>Loading...</div>;
	}
	return (
		<IngredientCreateFormView ingredientCategoryList={ingredientCategoryList} />
	);
};
