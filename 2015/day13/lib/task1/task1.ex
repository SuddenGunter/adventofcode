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


  defp permutations([]), do: [[]]

  defp permutations(list),
    do: for(elem <- list, rest <- permutations(list -- [elem]), do: [elem | rest])

  # defp generatePermutations(arrangement, len) when len == 1 do
  #   Enum.map(arrangement, fn x -> [x] end) |> IO.inspect()
  # end

  # defp generatePermutations(arrangement, len) do
  #   generatePermutations(arrangement, len - 1)
  #   |> Enum.flat_map(fn t ->
  #     rest = Enum.filter(
  #       arrangement,
  #       fn e ->
  #         not Enum.any?(t, fn j -> j == e end)
  #       end
  #     )
  #     [t | [rest]]
  #   end)
  # end

  # defp generatePermutations(arrangement) do
  #   generatePermutations(arrangement, length(arrangement))
  # end

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
