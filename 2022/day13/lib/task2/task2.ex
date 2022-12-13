defmodule Task2 do
  @spec solution([{[], []}]) :: integer()
  def solution(pairs) do
    Enum.reduce(pairs, [], fn {l, r}, acc -> [l] ++ [r] ++ acc end)
    |> with_new_pairs()
    |> Enum.sort(&Comparator.compare/2)
    |> Enum.with_index()
    |> Enum.filter(fn {x, _} -> x == [[2]] or x == [[6]] end)
    |> Enum.map(fn {_, i} -> i + 1 end)
    |> Enum.reduce(1, fn x, acc -> x * acc end)
  end

  @spec with_new_pairs([]) :: []
  defp with_new_pairs(list) do
    [[[2]]] ++ [[[6]]] ++ list
  end
end
