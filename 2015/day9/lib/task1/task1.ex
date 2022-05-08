defmodule Task1 do
  @spec solution(%{}) :: integer()
  def solution(graph) do
    cities = graph |> Enum.map(fn {k, _} -> elem(k, 0) end) |> Enum.dedup()

    paths =
      cities
      |> Enum.map(fn x -> solveForCity(graph, cities, x, [x], 0, 0) end)
      |> IO.inspect(label: "all paths")
      |> Enum.filter(fn x -> elem(x, 0) != :no_path end)

    case length(paths) do
      0 ->
        0

      _ ->
        Enum.min_by(paths, fn {_, _, len} -> len end)
        |> IO.inspect(label: "min path")
        |> elem(2)
    end
  end

  defp solveForCity(graph, cities, from, visited, skipped, pathLen) do
    closest = closestNotVisited(graph, visited, skipped, from)

    cond do
      closest == :no_path and length(visited) == length(cities) ->
        {
          :ok,
          visited,
          pathLen
        }

      # todo: try next shortest
      closest == :no_path and length(visited) + skipped < length(cities) ->
        solveForCity(
          graph,
          cities,
          from,
          visited,
          skipped + 1,
          pathLen
        )

      closest == :no_path and length(visited) + skipped == length(cities) ->
        :no_path

      true ->
        solveForCity(
          graph,
          cities,
          closest |> elem(0) |> elem(1),
          [closest |> elem(0) |> elem(1) | visited],
          skipped,
          pathLen + elem(closest, 1)
        )
    end
  end

  defp closestNotVisited(graph, visited, skipped, from) do
    availablePaths = graph |> Enum.filter(filter(visited, from))

    case availablePaths do
      [_ | _] ->
        availablePaths
        |> Enum.sort_by(fn {_, v} -> v end)
        |> Enum.drop(skipped)
        |> Enum.min_by(fn {_, v} -> v end)

      [] ->
        :no_path
    end
  end

  defp filter(visited, from) do
    fn {k, _} -> elem(k, 0) == from and !Enum.any?(visited, fn x -> x == elem(k, 1) end) end
  end
end
