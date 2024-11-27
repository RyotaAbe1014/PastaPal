

import { useIngredients } from "../../hooks";
import { IngredientListTableView } from "./view";

export const IngredientListTable = () => {
	const { ingredients, isLoading } = useIngredients();

	if (isLoading || !ingredients) {
		return <div>Loading...</div>;
	}

	return (
		<IngredientListTableView ingredients={ingredients} />
	);
};
