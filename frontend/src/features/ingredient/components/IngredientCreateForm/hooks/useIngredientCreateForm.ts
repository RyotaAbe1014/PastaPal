import { ApiClient, isApiError } from "@/api/client";
import type { CreateIngredientRequest } from "@/api/types/createIngredientRequest";
import type { CreateIngredientResponse } from "@/api/types/createIngredientResponse";
import { useGetIngredientCategoryList, useGetIngredientList } from "@/features/ingredient/hooks";
import { useState } from "react";

export const useIngredientCreateForm = () => {
	const api = ApiClient();
	const [name, setName] = useState<string>("");
	const [ingredientCategoryId, setIngredientCategoryId] = useState<string>("");

	const { ingredientCategories, isLoading } = useGetIngredientCategoryList();
	const { mutate } = useGetIngredientList();

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

		mutate();
	};

	return {
		ingredientCategoryList: ingredientCategories,
		isLoading,
		name,
		setName,
		onConfirm,
		ingredientCategoryId,
		setIngredientCategoryId,
	};
};
