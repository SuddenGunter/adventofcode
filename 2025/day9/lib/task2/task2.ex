defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    red_tiles = parse(input)

    red_tiles
    |> area_heap()
    |> find_largest_allowed(polygon(red_tiles))
  end

  # in task 2, we can only take rectangles inside outer polygon
  defp find_largest_allowed(heap, polygon) do
    {area, {x1, y1}, {x2, y2}} = Heap.root(heap)

    case contains?(polygon, polygon([{x1, y1}, {x2, y1}, {x2, y2}, {x1, y2}])) do
      true ->
        area

      false ->
        find_largest_allowed(Heap.pop(heap), polygon)
    end
  end

  defp polygon(points) do
    %Geo.Polygon{coordinates: [points]}
  end

  defp contains?(outer, inner) do
    Topo.contains?(outer, inner)
  end

  defp parse(input) do
    String.trim(input)
    |> String.split("\n")
    |> Enum.map(fn line ->
      [col, row] = String.split(line, ",")
      {parse_int!(row), parse_int!(col)}
    end)
  end

  defp area_heap(tiles) do
    Enum.reduce(tiles, {Heap.new(:>), Enum.drop(tiles, 1)}, fn tile, {acc, following_tiles} ->
      {Enum.map(following_tiles, fn x ->
         # {area, from, to}
         {area(tile, x), tile, x}
       end)
       |> Enum.reduce(acc, fn x, acc_heap -> Heap.push(acc_heap, x) end),
       Enum.drop(following_tiles, 1)}
    end)
    |> elem(0)
  end

  def area({x1, y1}, {x2, y2}) do
    (1 + abs(x1 - x2)) * (1 + abs(y1 - y2))
  end

  defp parse_int!(num) do
    {x, _} = Integer.parse(num)
    x
  end
end
