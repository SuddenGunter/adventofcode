defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    input
    |> parse_input()
    |> process()
  end

  defp process({rows, cols, map}) do
    map
    |> Map.keys()
    |> Enum.flat_map(fn x ->
      antenas = Map.get(map, x)
      pairs = for first <- antenas, second <- antenas, first != second, do: {first, second}

      pairs
      |> Enum.map(fn {{lx, ly}, {rx, ry}} ->
        {vecX, vecY} = {lx - rx, ly - ry}
        {lx + vecX, ly + vecY}
      end)
    end)
    |> Enum.filter(fn {x, y} ->
      cond do
        x < 0 or y < 0 -> false
        x > rows || y > cols -> false
        true -> true
      end
    end)
    |> Enum.uniq()
    |> length()
  end

  defp parse_input(input) do
    {rows, cols, map} =
      input
      |> String.trim()
      |> String.to_charlist()
      |> Enum.reduce(
        {0, 0, %{}},
        fn
          ?\n, {row, _, state} ->
            {row + 1, 0, state}

          letter, {row, col, state} ->
            {row, col + 1, upsert_mapset(state, {row, col}, letter)}
        end
      )

    # -1 cause \n affect row length
    {rows, cols - 1, map |> Map.delete(?.)}
  end

  defp upsert_mapset(state, {row, col}, letter) do
    case Map.get(state, letter) do
      nil -> Map.put(state, letter, MapSet.new() |> MapSet.put({row, col}))
      ms -> %{state | letter => MapSet.put(ms, {row, col})}
    end
  end
end
