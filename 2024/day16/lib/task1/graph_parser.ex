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
          ?#, {row, col} -> {:skip, {row, col + 1}}
          x, {row, col} -> {{row, col, x}, {row, col + 1}}
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

    nodes_values = Enum.reduce(nodes, %{}, fn {r, c, x}, m -> Map.put(m, {r, c}, x) end)

    max_size =
      Enum.reduce(nodes_values, {0, 0}, fn {{r, c}, _}, {mr, mc} -> {max(r, mr), max(c, mc)} end)

    to_add_edges =
      Graph.vertices(vertexes)
      |> Enum.reduce([], fn v, acc ->
        neighbours = get_all_neighbours(v, max_size)

        edges =
          Enum.map(neighbours, fn {row, col} ->
            # poor-mans undirected graph using directed graph.
            # libgraph undirected graph support is broken, some algorithms does not work
            # so using next best thing
            Graph.Edge.new(v, {row, col})
            Graph.Edge.new({row, col}, v)
          end)

        edges ++ acc
      end)
      |> Enum.filter(fn %{v1: v1, v2: v2} ->
        Map.has_key?(nodes_values, v1) and Map.has_key?(nodes_values, v2)
      end)

    final_graph = to_add_edges |> Enum.reduce(vertexes, fn edge, g -> Graph.add_edge(g, edge) end)

    start =
      Enum.find(nodes_values, fn {_k, v} ->
        v == ?S
      end)
      |> elem(0)

    finish =
      Enum.find(nodes_values, fn {_k, v} ->
        v == ?E
      end)
      |> elem(0)

    %{g: final_graph, start: start, finish: finish}
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
