import { GetIngredientListResponse } from "@/api/types/getIngredientListResponse";
import { fetcher } from "@/libs/swr/fetcher"
import useSWR from "swr"

export const useIngredients = () => {
  const { data, isLoading, error, isValidating, mutate } = useSWR<GetIngredientListResponse>("/ingredients", fetcher<undefined, GetIngredientListResponse>);

  if (error) {
    console.error(error);
    // TODO: エラー処理
  }
  return {
    ingredients: data?.data,
    isLoading,
    isValidating,
    mutate,
    error,
  }
}