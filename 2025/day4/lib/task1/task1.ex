defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    parse(input)
    |> filter_by_adjacent_rolls(4)
    |> length()
  end

  def parse(input) do
    String.trim(input)
    |> String.split("\n")
    |> Enum.with_index()
    |> Enum.reduce(%{}, fn {line, line_num}, roll_positions ->
      String.graphemes(line)
      |> Enum.with_index()
      |> Enum.reduce(roll_positions, fn {x, i}, acc ->
        Map.put(acc, {line_num, i}, x)
      end)
    end)
  end

  defp filter_by_adjacent_rolls(rolls, limit) do
    # we only care about rolls positions, not empty slots
    Enum.filter(
      rolls,
      fn
        {_pos, "."} -> false
        {_pos, "@"} -> true
      end
    )
    |> Enum.filter(fn {{x, y}, "@"} ->
      adjacent =
        Enum.filter(
          rolls,
          fn
            {{i, j}, "@"}
            when i in [x - 1, x, x + 1] and j in [y - 1, y, y + 1] and not ({x, y} == {i, j}) ->
              true

            _ ->
              false
          end
        )
        |> length()

      adjacent < limit
    end)
  end
end
