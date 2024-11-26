import type { GetIngredientCategoryListResponse } from "@/api/types/getIngredientCategoryListResponse";
import { fetcher } from "@/libs/swr/fetcher";
import useSWR from "swr";

export const useIngredientCreateForm = () => {
	const { data, isLoading, error } = useSWR<GetIngredientCategoryListResponse>(
		"/ingredient-categories",
		(url: string) => fetcher<undefined, GetIngredientCategoryListResponse>(url),
	);

	if (error) {
		console.error(error);
		// TODO: エラーをダイアログで表示する
	}

	return {
		ingredientCategoryList: data?.data,
		isLoading,
	};
};
