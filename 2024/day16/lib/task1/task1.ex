defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    GraphParser.parse(input)
    |> find_shortest_path()
  end

  defp find_shortest_path(%{g: g, start: s, finish: f}) do
    paths = Graph.Pathfinding.all(g, s, f)
    IO.inspect(length(paths), label: :pathlen)
    paths
    |> Enum.map(fn x ->
      get_path_cost(x)
    end)
    |> Enum.min()
  end

  def get_path_cost([start | tail]) do
    # east
    start_direction = {0, 1}
    path_cost(start_direction, start, tail, 0)
  end

  defp path_cost(_, _, [], acc) do
    acc
  end

  defp path_cost(direction, prev, [v | tail], acc) do
    new_direction = diff(v, prev)

    if new_direction == direction do
      path_cost(direction, v, tail, acc + 1)
    else
      if valid_rotation?(direction, new_direction) do
        path_cost(new_direction, v, tail, acc + 1000 + 1)
      else
        # atoms are always bigger then any integer
        :infinity
      end
    end
  end

  def valid_rotation?({x, y}, {x1, y1}) do
    # can only rotate 90 degrees
    abs(x) + abs(x1) <= 1 and abs(y) + abs(y1) == 1
  end

  defp diff({x, y}, {x1, y1}) do
    {x - x1, y - y1}
  end
end
