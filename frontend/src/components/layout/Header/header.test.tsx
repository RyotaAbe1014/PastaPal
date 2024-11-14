import { Wrapper } from "@/libs/test/wrapper";
import { ChakraProvider, defaultSystem } from "@chakra-ui/react";
import { render, screen } from "@testing-library/react";
import { beforeEach, describe, expect, it } from "vitest";
import { vi } from "vitest";
import { Header } from "./header";

vi.mock("@tanstack/react-router", () => ({
	Link: ({ children, to }: { children: React.ReactNode; to: string }) => (
		<a href={to}>{children}</a>
	),
	useNavigate: () => vi.fn(),
}));

beforeEach(() => {
	vi.resetAllMocks();
});

describe("Header Component", () => {
	it("renders the Header component with dashboard path", () => {
		render(
			<ChakraProvider value={defaultSystem}>
				<Header path="/recipes" />
			</ChakraProvider>,
			{ wrapper: Wrapper },
		);
		expect(screen.getByText("ダッシュボード")).toBeInTheDocument();
		expect(screen.getByText("レシピ")).toBeInTheDocument();
		expect(screen.getByText("食材")).toBeInTheDocument();
	});
});
