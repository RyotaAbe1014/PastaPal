import { server } from "@/libs/test/server";
import { act, renderHook } from "@testing-library/react";
import { http, HttpResponse } from "msw";
import { beforeEach, describe, expect, it, vi } from "vitest";
import { useIngredientCategoryCreateForm } from "./useIngredientCategoryCreateForm";

beforeEach(() => {
	server.resetHandlers();
	vi.resetAllMocks();
	server.use(
		http.get("http://localhost:8000/api/v1/ingredient-categories", () => {
			return HttpResponse.json({ data: [] });
		}),
	);
});

describe("useIngredientCategoryCreateForm", () => {
	it("初期値が正しく設定されていること", () => {
		const { result } = renderHook(() => useIngredientCategoryCreateForm());

		expect(result.current.name).toBe("");
	});

	it("setNameが正しく動作すること", () => {
		const { result } = renderHook(() => useIngredientCategoryCreateForm());

		act(() => {
			result.current.setName("新しいカテゴリー");
		});

		expect(result.current.name).toBe("新しいカテゴリー");
	});

	it("onConfirmが正しく動作すること", async () => {
		server.use(
			http.post("http://localhost:8000/api/v1/ingredient-categories", () => {
				return HttpResponse.json({ name: "新しいカテゴリー" });
			}),
		);
		const { result } = renderHook(() => useIngredientCategoryCreateForm());

		act(() => {
			result.current.setName("新しいカテゴリー");
		});

		await act(async () => {
			await result.current.onConfirm();
		});
	});
});
