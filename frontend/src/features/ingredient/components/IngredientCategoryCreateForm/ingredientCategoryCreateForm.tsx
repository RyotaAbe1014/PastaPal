import { useIngredientCategoryCreateForm } from "./hooks/useIngredientCategoryCreateForm";
import { IngredientCategoryCreateFormView } from "./view";

export const IngredientCategoryCreateForm = () => {
	const { name, setName, onConfirm } = useIngredientCategoryCreateForm();
	return <IngredientCategoryCreateFormView name={name} setName={setName} onConfirm={onConfirm} />;
};
