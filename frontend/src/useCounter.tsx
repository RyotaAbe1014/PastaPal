// render hooksパターンを使用したカウンター
import { atom, useAtom } from "jotai";

const countAtom = atom(0);

export const useCounter = () => {
  const [count, setCount] = useAtom(countAtom);

  const counter = (
    <div>
      <button onClick={() => setCount((c) => c + 1)}>count is {count}</button>
    </div>
  )

  return [
    counter,
  ]
}