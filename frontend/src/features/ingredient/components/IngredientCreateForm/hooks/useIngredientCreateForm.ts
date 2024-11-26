import { ApiClient, isApiError } from "@/api/client";
import type { CreateIngredientRequest } from "@/api/types/createIngredientRequest";
import type { CreateIngredientResponse } from "@/api/types/createIngredientResponse";
import type { GetIngredientCategoryListResponse } from "@/api/types/getIngredientCategoryListResponse";
import { fetcher } from "@/libs/swr/fetcher";
import { useState } from "react";
import useSWR from "swr";

export const useIngredientCreateForm = () => {
	const api = ApiClient();
	const [name, setName] = useState<string>("");
	const [ingredientCategoryId, setIngredientCategoryId] = useState<string>("");

	const { data, isLoading, error } = useSWR<GetIngredientCategoryListResponse>(
		"/ingredient-categories",
		(url: string) => fetcher<undefined, GetIngredientCategoryListResponse>(url),
	);

	if (error) {
		console.error(error);
		// TODO: エラーをダイアログで表示する
	}

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
		ingredientCategoryList: data?.data,
		isLoading,
		name,
		setName,
		onConfirm,
		ingredientCategoryId,
		setIngredientCategoryId,
	};
};
