defmodule Task2 do
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
      |> Enum.flat_map(fn {{lx, ly}, {rx, ry}} ->
        {vecX, vecY} = {lx - rx, ly - ry}

        [{lx, ly}, {rx, ry}] ++
          generate_antinodes({lx, ly}, {vecX, vecY}, {rows, cols})
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

  defp generate_antinodes({x, y}, {vecX, vecY}, {rows, cols}) do
    {newX, newY} = {x + vecX, y + vecY}

    valid =
      cond do
        newX < 0 or newY < 0 -> false
        newX > rows || newY > cols -> false
        true -> true
      end

    if valid do
      [{newX, newY}] ++ generate_antinodes({newX, newY}, {vecX, vecY}, {rows, cols})
    else
      []
    end
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
