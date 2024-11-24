import { ApiClient } from "@/api/client";
import type { getIngredientCategoryListResponse } from "@/api/types/getIngredientCategoryListResponse";
import useSWR from "swr";

export const useIngredientCreateForm = () => {
	const api = ApiClient();

	const { data, isLoading } = useSWR<getIngredientCategoryListResponse>(
		"/ingredient-categories",
		(url: string) => api.Get<undefined, getIngredientCategoryListResponse>(url),
	);

	return {
		ingredientCategoryList: data?.data,
		isLoading,
	};
};
