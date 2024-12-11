defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    {g, all_possible_interesting_routes} = GraphParser.parse(input)

    all_possible_interesting_routes
    |> Enum.flat_map(fn {from, to} ->
      Graph.get_paths(g, from, to)
    end)
    |> Enum.count()
  end
end
