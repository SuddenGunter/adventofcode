defmodule Task1 do
  @spec solution(%{}) :: integer()
  def solution(costs) do
    default = costs |> defaultArrangement
    first = List.first(default)

    arrangements =
      Enum.drop(default, 1)
      |> permutations
      |> Enum.map(fn x -> [first | x] end)

    Enum.map(arrangements, fn x -> cost(x, costs) end)
    |> Enum.max()
  end

  def permutations([]), do: [[]]

  def permutations(list),
    do: for(elem <- list, rest <- permutations(list -- [elem]), do: [elem | rest])

  defp cost(arrangement, costs) do
    sum =
      Enum.zip(arrangement, Enum.drop(arrangement, 1))
      |> Enum.reduce(0, fn {l, r}, acc ->
        acc + pair_cost(l, r, costs)
      end)

    sum + pair_cost(List.first(arrangement), List.last(arrangement), costs)
  end

  # todo: memoize?
  defp pair_cost(l, r, costs) do
    costs[l].costs[r] + costs[r].costs[l]
  end

  defp defaultArrangement(costs) do
    Enum.map(costs, fn {k, _} -> k end)
  end
end
