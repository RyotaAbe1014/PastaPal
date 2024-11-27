import { ApiClient, isApiError } from "@/api/client";
import type { CreateIngredientRequest } from "@/api/types/createIngredientRequest";
import type { CreateIngredientResponse } from "@/api/types/createIngredientResponse";
import { useGetIngredientCategoryList } from "@/features/ingredient/hooks";
import { useState } from "react";

export const useIngredientCreateForm = () => {
	const api = ApiClient();
	const [name, setName] = useState<string>("");
	const [ingredientCategoryId, setIngredientCategoryId] = useState<string>("");

	const { ingredients, isLoading } = useGetIngredientCategoryList();

	const onConfirm = async () => {
		if (!ingredientCategoryId) {
			// TODO: エラーをダイアログで表示する
			return;
		}

		const response = await api.Post<
			CreateIngredientRequest,
			CreateIngredientResponse
		>("/ingredients", {
			name,
			ingredientCategoryId: ingredientCategoryId,
		});

		if (isApiError(response)) {
			console.error(response.message);
			// TODO: エラーをダイアログで表示する
			return;
		}
	};

	return {
		ingredientCategoryList: ingredients,
		isLoading,
		name,
		setName,
		onConfirm,
		ingredientCategoryId,
		setIngredientCategoryId,
	};
};
