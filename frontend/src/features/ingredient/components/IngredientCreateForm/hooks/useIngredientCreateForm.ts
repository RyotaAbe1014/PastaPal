import { ApiClient, type ApiResponse } from "@/api/client";
import type { GetIngredientCategoryListResponse } from "@/api/types/getIngredientCategoryListResponse";
import useSWR from "swr";

export const useIngredientCreateForm = () => {
	const api = ApiClient();

	const { data, isLoading } = useSWR<
		ApiResponse<GetIngredientCategoryListResponse>
	>("/ingredient-categories", (url: string) =>
		api.Get<undefined, GetIngredientCategoryListResponse>(url),
	);

	return {
		ingredientCategoryList: data?.data,
		isLoading,
	};
};
