import { ApiClient } from "@/api/client";
import type { CreateIngredientCategoryRequest } from "@/api/types/createIngredientCategoryRequest";
import type { CreateIngredientCategoryResponse } from "@/api/types/createIngredientCategoryResponse";
import { useState } from "react";

export const useIngredientCategoryCreateForm = () => {
	const api = ApiClient();
	const [name, setName] = useState("");

	const onConfirm = async () => {
		await api.Post<
			CreateIngredientCategoryRequest,
			CreateIngredientCategoryResponse
		>("/ingredient-categories", { name });
	};

	return {
		name,
		setName,
		onConfirm,
	};
};
