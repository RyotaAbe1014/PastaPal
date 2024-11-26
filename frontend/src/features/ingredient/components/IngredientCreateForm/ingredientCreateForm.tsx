import { useIngredientCreateForm } from "./hooks";
import { IngredientCreateFormView } from "./views";

export const IngredientCreateForm = () => {
	const {
		ingredientCategoryList,
		isLoading,
		name,
		setName,
		ingredientCategoryId,
		setIngredientCategoryId,
		onConfirm,
	} = useIngredientCreateForm();

	if (isLoading || typeof ingredientCategoryList === "undefined") {
		return <div>Loading...</div>;
	}
	return (
		<IngredientCreateFormView
			ingredientCategoryList={ingredientCategoryList}
			name={name}
			setName={setName}
			ingredientCategoryId={ingredientCategoryId}
			setIngredientCategoryId={setIngredientCategoryId}
			onConfirm={onConfirm}
		/>
	);
};
