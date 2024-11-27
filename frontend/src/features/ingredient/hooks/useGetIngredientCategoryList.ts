import { GetIngredientCategoryListResponse } from "@/api/types/getIngredientCategoryListResponse";
import { fetcher } from "@/libs/swr/fetcher";
import useSWR from "swr";

export const useGetIngredientCategoryList = () => {
  const { data, isLoading, error, isValidating, mutate } = useSWR<GetIngredientCategoryListResponse>(
    "/ingredient-categories",
    (url: string) => fetcher<undefined, GetIngredientCategoryListResponse>(url),
  );

  if (error) {
    console.error(error);
    // TODO: エラーをダイアログで表示する
  }

  return {
    ingredients: data?.data,
    isLoading,
    isValidating,
    mutate,
    error,
  };
};
