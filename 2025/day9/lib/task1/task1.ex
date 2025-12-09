defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    parse(input)
    |> distance_heap()
    |> Heap.root()
    |> area()
  end

  defp parse(input) do
    String.trim(input)
    |> String.split("\n")
    |> Enum.map(fn line ->
      [col, row] = String.split(line, ",")
      {parse_int!(row), parse_int!(col)}
    end)
  end

  defp distance_heap(tiles) do
    Enum.reduce(tiles, {Heap.new(:>), Enum.drop(tiles, 1)}, fn tile, {acc, following_tiles} ->
      {Enum.map(following_tiles, fn x ->
         # {distance, from, to}
         {manhattan_distance(tile, x), tile, x}
       end)
       |> Enum.reduce(acc, fn x, acc_heap -> Heap.push(acc_heap, x) end),
       Enum.drop(following_tiles, 1)}
    end)
    |> elem(0)
  end

  defp manhattan_distance({x1, y1}, {x2, y2}) do
    1 + abs(x1 - x2) + (1 + abs(y1 - y2))
  end

  def area({_, {x1, y1}, {x2, y2}}) do
    (1 + abs(x1 - x2)) * (1 + abs(y1 - y2))
  end

  defp parse_int!(num) do
    {x, _} = Integer.parse(num)
    x
  end
end
