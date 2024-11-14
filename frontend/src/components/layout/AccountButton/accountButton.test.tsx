import { fireEvent, render, screen } from "@testing-library/react";
import { describe, it, vi, expect, beforeEach } from "vitest";
import { AccountButton } from "./accountButton";
import { Wrapper } from "@/libs/test/wrapper";
import { act } from "react";

vi.mock("@/api/client");
vi.mock("@tanstack/react-router", () => ({
  useNavigate: vi.fn(),
}));

describe("AccountButton", () => {
  beforeEach(() => {
    // localStorageのモックを作成
    const localStorageMock = (function () {
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

    Object.defineProperty(window, 'localStorage', {
      value: localStorageMock,
    });
  });
  it("アカウントボタンが表示されること", () => {
    render(
      <AccountButton />,
      { wrapper: Wrapper },
    );
    expect(screen.getByLabelText("アカウント")).toBeInTheDocument();
  });

  it("ユーザー名が表示されること", async () => {
    localStorage.setItem("userId", "testUser");
    render(
      <AccountButton />,
      { wrapper: Wrapper },
    );

    await act(async () => {
      fireEvent.click(screen.getByLabelText('アカウント'));
    });
      expect(screen.getByText("ユーザー名: testUser")).toBeInTheDocument();
  });

  // it("logs out the user when the logout button is clicked", async () => {
  //   const mockNavigate = vi.fn();
  //   const mockPost = vi.fn().mockResolvedValue({});
  //   (useNavigate as vi.Mock).mockReturnValue(mockNavigate);
  //   (ApiClient as vi.Mock).mockReturnValue({ Post: mockPost });

  //   render(<AccountButton />);
  //   fireEvent.click(screen.getByLabelText("アカウント"));
  //   fireEvent.click(screen.getByText("ログアウト"));

  //   expect(localStorage.getItem("userId")).toBeNull();
  //   expect(mockPost).toHaveBeenCalledWith("/auth/github/logout");
  //   expect(mockNavigate).toHaveBeenCalledWith({ to: "/login" });
  // });
});