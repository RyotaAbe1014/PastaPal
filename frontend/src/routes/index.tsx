import { Button } from "@/components/ui/button";
import { createFileRoute } from "@tanstack/react-router";
import { atom, useAtom } from "jotai";
import useSWR from "swr";

export const Route = createFileRoute("/")({
	component: Index,
});

const countAtom = atom(0);

export function Index() {
	const [count, setCount] = useAtom(countAtom);

	const { isLoading, isValidating, data } = useSWR(
		"http://localhost:8000/health",
		async (url) => {
			const res = await fetch(url);
			return res;
		},
	);

	return (
		<>
			<h1>Vite + React</h1>
			<div className="card">
				<Button onClick={() => setCount((count) => count + 1)}>
					count is {count}
				</Button>
				<p>
					Edit <code>src/App.tsx</code> and save to test HMR
				</p>
			</div>
			<div className="card">
				using Jotai
				{/* {counter} */}
			</div>

			<div className="card">
				<h2>SWR</h2>
				{isLoading ? (
					<div>Loading...</div>
				) : isValidating ? (
					<div>Validating...</div>
				) : (
					<div>{data ? "Success" : "Failed"}</div>
				)}
			</div>
		</>
	);
}
