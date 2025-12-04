defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    parse(input)
    |> Enum.filter(fn
      {_pos, "."} -> false
      {_pos, "@"} -> true
    end)
    |> remove_while_possible(0)
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

  defp remove_while_possible(rolls, removed) do
    case filter_by_adjacent_rolls(rolls, 4) do
      [] ->
        removed

      to_remove ->
        remove(rolls, to_remove) |> remove_while_possible(removed + length(to_remove))
    end
  end

  defp remove(rolls, to_remove) do
    removable_coordinates = Enum.map(to_remove, fn {x, "@"} -> x end)

    Enum.filter(rolls, fn {{x, y}, "@"} ->
      !Enum.any?(removable_coordinates, fn c -> c == {x, y} end)
    end)
  end

  defp filter_by_adjacent_rolls(rolls, limit) do
    Enum.filter(rolls, fn {{x, y}, "@"} ->
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
