import { RouterProvider, createRouter } from "@tanstack/react-router";
import { StrictMode } from "react";
import ReactDOM from "react-dom/client";

import { SWRConfig } from "swr";
import { Provider } from "./components/ui/provider";
// Import the generated route tree
import { routeTree } from "./routeTree.gen";

// Create a new router instance
const router = createRouter({ routeTree });

// Register the router instance for type safety
declare module "@tanstack/react-router" {
	interface Register {
		router: typeof router;
	}
}

// Render the app
const rootElement = document.getElementById("root");
if (rootElement && !rootElement.innerHTML) {
	const root = ReactDOM.createRoot(rootElement);
	root.render(
		<StrictMode>
			<Provider>
				<SWRConfig
					value={{
						revalidateOnFocus: false,
						revalidateOnReconnect: false,
					}}
				>
					<RouterProvider router={router} />
				</SWRConfig>
			</Provider>
		</StrictMode>,
	);
}
