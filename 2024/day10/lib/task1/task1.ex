defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    {g, all_possible_interesting_routes} = GraphParser.parse(input)

    all_possible_interesting_routes
    |> Enum.filter(fn {from, to} ->
      Graph.reachable(g, [from])
      |> Enum.member?(to)
    end)
    |> Enum.count()
  end
end
