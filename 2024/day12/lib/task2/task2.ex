defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    g = GraphParser.parse(input)

    Graph.components(g)
    |> Enum.map(fn subgraph ->
      area = length(subgraph)
      sides = find_corners(g, subgraph)
      area * sides
    end)
    |> Enum.sum()
  end

  def find_corners(g, subgraph) do
    Enum.map(subgraph, fn v ->
      [
        external_corner?(g, v),
        internal_corner?(g, subgraph, v)
      ]
      |> Enum.flat_map(fn x -> x end)
      |> Enum.filter(fn x -> x end)
      |> Enum.count()
    end)
    |> Enum.sum()
  end

  def internal_corner?(g, subgraph, {row, col}) do
    n = Graph.neighbors(g, {row, col})

    [
      not Enum.member?(subgraph, {row + 1, col + 1}) and Enum.member?(n, {row + 1, col}) and
        Enum.member?(n, {row, col + 1}),
      not Enum.member?(subgraph, {row - 1, col - 1}) and Enum.member?(n, {row - 1, col}) and
        Enum.member?(n, {row, col - 1}),
      not Enum.member?(subgraph, {row + 1, col - 1}) and Enum.member?(n, {row + 1, col}) and
        Enum.member?(n, {row, col - 1}),
      not Enum.member?(subgraph, {row - 1, col + 1}) and Enum.member?(n, {row - 1, col}) and
        Enum.member?(n, {row, col + 1})
    ]
  end

  def external_corner?(g, {row, col}) do
    n = Graph.neighbors(g, {row, col})

    [
      not Enum.member?(n, {row + 1, col}) and not Enum.member?(n, {row, col + 1}),
      not Enum.member?(n, {row - 1, col}) and not Enum.member?(n, {row, col - 1}),
      not Enum.member?(n, {row + 1, col}) and not Enum.member?(n, {row, col - 1}),
      not Enum.member?(n, {row - 1, col}) and not Enum.member?(n, {row, col + 1})
    ]
  end
end

# ...
# ##.
# ##.

# ...
# .##
# .##

# ##.
# ##.
# ...

# .##
# .##
# ...
