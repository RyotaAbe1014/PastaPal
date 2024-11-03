import { ChakraProvider, defaultSystem } from "@chakra-ui/react";
import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { Header } from "./header";

describe("Header Component", () => {
	it("renders the Header component with dashboard path", () => {
		render(
			<ChakraProvider value={defaultSystem}>
				<Header path="dashboard" />
			</ChakraProvider>,
		);
		expect(screen.getByText("ダッシュボード")).toBeInTheDocument();
		expect(screen.getByText("レシピ")).toBeInTheDocument();
		expect(screen.getByText("食材")).toBeInTheDocument();
	});
});
