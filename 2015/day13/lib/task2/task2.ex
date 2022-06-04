defmodule Task2 do
  @spec solution(%{}) :: integer()
  def solution(costs) do
    costsWithMe = addMe(costs)
    default = costsWithMe |> defaultArrangement
    first = List.first(default)

    arrangements =
      Enum.drop(default, 1)
      |> permutations
      |> Enum.map(fn x -> [first | x] end)

    Enum.map(arrangements, fn x -> cost(x, costsWithMe) end)
    |> Enum.max()
  end

  defp addMe(costs) do
    allWithMe =
      defaultArrangement(costs)
      |> Enum.reduce(costs, fn x, acc ->
        Map.put(acc, x, %Person{costs: Map.put(costs[x].costs, "me", 0)})
      end)

    Map.put(allWithMe, "me", %Person{
      costs:
        defaultArrangement(costs)
        |> Enum.reduce(%{}, fn x, acc ->
          Map.put(acc, x, 0)
        end)
    })
  end

  defp permutations([]), do: [[]]

  defp permutations(list),
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
