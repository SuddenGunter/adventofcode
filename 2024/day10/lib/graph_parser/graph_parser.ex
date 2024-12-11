defmodule GraphParser do
  # parse graph from input and return graph with only valid edges (both for p1 and p2) and also
  # return all possible interesting trail starts+finish combinations
  @spec parse(String.t()) :: {Graph.t(), [{{integer(), integer()}, {integer(), integer()}}]}
  def parse(input) do
    nodes =
      input
      |> String.trim()
      |> String.to_charlist()
      |> Enum.map_reduce(
        {0, 0},
        fn
          ?\n, {row, _col} -> {:skip, {row + 1, 0}}
          x, {row, col} -> {{row, col, x - ?0}, {row, col + 1}}
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
        Graph.new(),
        fn {row, col, _}, g -> Graph.add_vertex(g, {row, col}) end
      )

    nodes_weights = Enum.reduce(nodes, %{}, fn {r, c, x}, m -> Map.put(m, {r, c}, x) end)
    max_size = nodes_weights |> Map.keys() |> Enum.sort(:desc) |> Enum.at(0)

    to_add_edges =
      Graph.vertices(vertexes)
      |> Enum.reduce([], fn v, acc ->
        weight = Map.get(nodes_weights, v)
        neighbours = get_all_neighbours(v, max_size)

        edges =
          Enum.map(neighbours, fn {row, col} ->
            edge_w = Map.get(nodes_weights, {row, col}) - weight
            Graph.Edge.new(v, {row, col}, weight: edge_w)
          end)

        edges ++ acc
      end)
      |> Enum.filter(fn %{weight: w} ->
        # we only care about edges that increase value by 1
        # we can't go back (negative numbers), we can't go higher then 1, we can't stay on same level
        if w != 1 do
          false
        else
          true
        end
      end)

    g = to_add_edges |> Enum.reduce(vertexes, fn edge, g -> Graph.add_edge(g, edge) end)

    {zeroes, nines} =
      nodes_weights
      |> Enum.reduce({[], []}, fn {{row, col}, w}, {zeroes, nines} ->
        case w do
          0 -> {[{row, col}] ++ zeroes, nines}
          9 -> {zeroes, [{row, col}] ++ nines}
          _ -> {zeroes, nines}
        end
      end)

    all_possible_interesting_routes = for x <- zeroes, y <- nines, do: {x, y}

    {g, all_possible_interesting_routes}
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
