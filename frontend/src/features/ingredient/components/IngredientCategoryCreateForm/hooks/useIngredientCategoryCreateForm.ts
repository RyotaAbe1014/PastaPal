import { ApiClient, isApiError } from "@/api/client";
import type { CreateIngredientCategoryRequest } from "@/api/types/createIngredientCategoryRequest";
import type { CreateIngredientCategoryResponse } from "@/api/types/createIngredientCategoryResponse";
import { useGetIngredientCategoryList } from "@/features/ingredient/hooks";
import { useState } from "react";

export const useIngredientCategoryCreateForm = () => {
	const api = ApiClient();
	const [name, setName] = useState("");
	const { mutate } = useGetIngredientCategoryList();

	const onConfirm = async () => {
		const response = await api.Post<
			CreateIngredientCategoryRequest,
			CreateIngredientCategoryResponse
		>("/ingredient-categories", { name });

		if (isApiError(response)) {
			// TODO: ダイアログ側にエラーを表示する
			return;
		}

		mutate();
		setName("");
	};

	return {
		name,
		setName,
		onConfirm,
	};
};
