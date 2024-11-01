// render hooksパターンを使用したカウンター
import { atom, useAtom } from "jotai";
import { Button } from "./components/ui/button";

const countAtom = atom(0);

export const useCounter = () => {
	const [count, setCount] = useAtom(countAtom);

	const counter = (
		<div>
			<Button type="button" onClick={() => setCount((c) => c + 1)}>
				count is {count}
			</Button>
		</div>
	);

	return [counter];
};
