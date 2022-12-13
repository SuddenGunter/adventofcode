defmodule Task1 do
  @spec solution([{[], []}]) :: integer()
  def solution(pairs) do
    Enum.map(pairs, fn {l, r} -> Comparator.compare(l, r) end)
    |> Enum.with_index()
    |> Enum.filter(fn {x, _} -> x end)
    |> Enum.map(fn {_, i} -> i + 1 end)
    |> Enum.sum()
  end
end
