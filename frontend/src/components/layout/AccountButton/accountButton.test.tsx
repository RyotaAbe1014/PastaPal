import { server } from "@/libs/test/server";
import { Wrapper } from "@/libs/test/wrapper";
import { useNavigate } from "@tanstack/react-router";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { http, HttpResponse } from "msw";
import { act } from "react";
import { type Mock, beforeEach, describe, expect, it, vi } from "vitest";
import { AccountButton } from "./accountButton";

vi.mock("@tanstack/react-router", () => ({
	useNavigate: vi.fn(),
}));

const mockedUseNavigate = useNavigate as Mock;

describe("AccountButton", () => {
	beforeEach(() => {
		// localStorageのモックを作成
		const localStorageMock = (() => {
			let store: { [key: string]: string } = {};
			return {
				getItem(key: string) {
					return store[key] || null;
				},
				setItem(key: string, value: string) {
					store[key] = value;
				},
				clear() {
					store = {};
				},
				removeItem(key: string) {
					delete store[key];
				},
			};
		})();

		Object.defineProperty(window, "localStorage", {
			value: localStorageMock,
		});
	});

	it("アカウントボタンが表示されること", () => {
		render(<AccountButton />, { wrapper: Wrapper });
		expect(screen.getByLabelText("アカウント")).toBeInTheDocument();
	});

	it("ユーザー名が表示されること", async () => {
		localStorage.setItem("userId", "testUser");
		render(<AccountButton />, { wrapper: Wrapper });

		await act(async () => {
			fireEvent.click(screen.getByLabelText("アカウント"));
		});
		expect(screen.getByText("ユーザー名: testUser")).toBeInTheDocument();
	});

	it("ログアウトボタンクリックでログアウト処理が実行されること", async () => {
		const mockedNavigate = vi.fn();
		mockedUseNavigate.mockReturnValue(mockedNavigate);
		server.use(
			http.post("http://localhost:8000/api/v1/auth/github/logout", () => {
				return HttpResponse.json({ isAuthenticated: false });
			}),
		);

		localStorage.setItem("userId", "testUser");

		render(<AccountButton />, { wrapper: Wrapper });

		await act(async () => {
			fireEvent.click(screen.getByLabelText("アカウント"));
		});

		await waitFor(() => {
			fireEvent.click(screen.getByRole("button", { name: "ログアウト" }));
		});

		expect(mockedNavigate).toHaveBeenCalledWith({ to: "/login" });
	});
});
