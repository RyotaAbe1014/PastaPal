import { ApiClient, isApiError } from "@/api/client";
import type { CreateIngredientCategoryRequest } from "@/api/types/createIngredientCategoryRequest";
import type { CreateIngredientCategoryResponse } from "@/api/types/createIngredientCategoryResponse";
import { useState } from "react";

export const useIngredientCategoryCreateForm = () => {
	const api = ApiClient();
	const [name, setName] = useState("");

	const onConfirm = async () => {
		const response = await api.Post<
			CreateIngredientCategoryRequest,
			CreateIngredientCategoryResponse
		>("/ingredient-categories", { name });

		if (isApiError(response)) {
			// TODO: ダイアログ側にエラーを表示する
			return;
		}

		setName("");
	};

	return {
		name,
		setName,
		onConfirm,
	};
};
