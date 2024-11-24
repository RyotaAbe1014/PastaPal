import { render, screen, fireEvent } from "@testing-library/react";
import { describe, it, vi, expect } from "vitest";
import { Wrapper } from "@/libs/test/wrapper";
import { IngredientCategoryCreateFormView } from "./ingredientCategoryCreateForm";

describe("IngredientCategoryCreateFormView", () => {
  it("初期値でフォームがレンダリングされること", () => {
    const mockOnConfirm = vi.fn();
    const mockSetName = vi.fn();

    render(
      <IngredientCategoryCreateFormView
        onConfirm={mockOnConfirm}
        name=""
        setName={mockSetName}
      />
      ,
      { wrapper: Wrapper }
    );

    expect(screen.getByPlaceholderText("新しい種別")).toBeInTheDocument();
    expect(screen.getByText("追加")).toBeInTheDocument();
    expect(screen.getByText("種別の追加")).toBeInTheDocument();
  });

  it("入力値が変更されたときにsetNameが呼び出されること", () => {
    const mockOnConfirm = vi.fn();
    const mockSetName = vi.fn();

    render(
      <IngredientCategoryCreateFormView
        onConfirm={mockOnConfirm}
        name=""
        setName={mockSetName}
      />
      ,
      { wrapper: Wrapper }
    );

    const input = screen.getByPlaceholderText("新しい種別");
    fireEvent.change(input, { target: { value: "新しいカテゴリー" } });

    expect(mockSetName).toHaveBeenCalledWith("新しいカテゴリー");
  });

  it("ボタンがクリックされたときにonConfirmが呼び出されること", () => {
    const mockOnConfirm = vi.fn();
    const mockSetName = vi.fn();

    render(
      <IngredientCategoryCreateFormView
        onConfirm={mockOnConfirm}
        name=""
        setName={mockSetName}
      />
      ,
      { wrapper: Wrapper }
    );

    const button = screen.getByText("追加");
    fireEvent.click(button);

    expect(mockOnConfirm).toHaveBeenCalled();
  });
});