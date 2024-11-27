import { useGetIngredientList } from "../../hooks";
import { IngredientListTableView } from "./view";

export const IngredientListTable = () => {
	const { ingredients, isLoading } = useGetIngredientList();

	if (isLoading || !ingredients) {
		return <div>Loading...</div>;
	}

	return <IngredientListTableView ingredients={ingredients} />;
};
