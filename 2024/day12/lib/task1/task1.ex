defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    g = GraphParser.parse(input)

    Graph.components(g)
    |> Enum.map(fn components ->
      area = length(components)

      perimeter =
        4 * area -
          (Enum.map(components, fn x ->
             Graph.neighbors(g, x) |> length()
           end)
           |> Enum.sum())

      area * perimeter
    end)
    |> Enum.sum()
  end
end
