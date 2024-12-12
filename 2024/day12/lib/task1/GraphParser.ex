defmodule GraphParser do
  @spec parse(String.t()) :: Graph.t()
  def parse(input) do
    nodes =
      input
      |> String.trim()
      |> String.to_charlist()
      |> Enum.map_reduce(
        {0, 0},
        fn
          ?\n, {row, _col} -> {:skip, {row + 1, 0}}
          x, {row, col} -> {{row, col, <<x::utf8>>}, {row, col + 1}}
        end
      )
      |> elem(0)
      |> Enum.filter(fn
        :skip -> false
        _ -> true
      end)

    vertexes =
      nodes
      |> Enum.reduce(
        Graph.new(type: :undirected),
        fn {row, col, _}, g -> Graph.add_vertex(g, {row, col}) end
      )

    nodes_values = Enum.reduce(nodes, %{}, fn {r, c, x}, m -> Map.put(m, {r, c}, x) end)

    {rows, cols, _} = nodes |> Enum.reverse() |> Enum.at(0)

    max_size = {rows, cols}

    to_add_edges =
      Graph.vertices(vertexes)
      |> Enum.reduce([], fn v, acc ->
        neighbours = get_all_neighbours(v, max_size)

        edges =
          Enum.map(neighbours, fn {row, col} ->
            Graph.Edge.new(v, {row, col})
          end)

        edges ++ acc
      end)
      |> Enum.filter(fn %{v1: v1, v2: v2} ->
        from = Map.get(nodes_values, v1)
        to = Map.get(nodes_values, v2)

        if from != to do
          false
        else
          true
        end
      end)

    to_add_edges |> Enum.reduce(vertexes, fn edge, g -> Graph.add_edge(g, edge) end)
  end

  defp get_all_neighbours({row, col}, {max_row, max_col}) do
    [
      {row - 1, col},
      {row, col - 1},
      {row + 1, col},
      {row, col + 1}
    ]
    |> Enum.filter(fn {k, v} ->
      # filter out neighbours outside of map borders
      k >= 0 and v >= 0 and k <= max_row and v <= max_col
    end)
  end
end
