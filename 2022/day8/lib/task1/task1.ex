defmodule Task1 do
  @spec solution([[integer()]]) :: integer()
  def solution(input) do
    rows_mins = rows_mins(input)

    transposed = transpose(input)

    col_mins = rows_mins(transposed) |> as_rows_mins()

    len_invisible =
      MapSet.intersection(rows_mins, col_mins)
      |> MapSet.to_list()
      |> Enum.filter(fn {row, col} -> row != 0 and col != 0 end)
      |> length()

    length(input) * length(transposed) - len_invisible
  end

  @spec rows_mins([[integer()]]) :: MapSet.t()
  defp rows_mins(matrix) do
    mapped =
      Enum.with_index(matrix, fn e, i -> {i, e} end)
      |> Enum.map(fn {row_ix, x} -> row_mins(Enum.drop(x, 1), 1, Enum.at(x, 0), row_ix) end)

    Enum.reduce(mapped, MapSet.new(), fn x, acc ->
      MapSet.union(x, acc)
    end)
  end

  @spec row_mins([integer()], pos_integer(), integer(), pos_integer()) :: MapSet.t()
  defp row_mins(row, dropped, known_max_left, row_ix) do
    if length(row) == 1 do
      MapSet.new()
    else
      cur_char = Enum.at(row, 0)
      new_known_max_left = max(known_max_left, cur_char)
      hidden_from_right = Enum.any?(Enum.drop(row, 1), fn x -> x >= cur_char end)
      next_row = row_mins(Enum.drop(row, 1), dropped + 1, new_known_max_left, row_ix)

      if hidden_from_right && cur_char <= known_max_left do
        MapSet.put(next_row, {row_ix, dropped})
      else
        next_row
      end
    end
  end

  @spec transpose([[integer()]]) :: [[integer()]]
  defp transpose(rows) do
    rows
    |> List.zip()
    |> Enum.map(&Tuple.to_list/1)
  end

  @spec as_rows_mins(MapSet.t()) :: MapSet.t()
  defp as_rows_mins(col_mins) do
    MapSet.to_list(col_mins)
    |> Enum.map(fn {r, c} -> {c, r} end)
    |> MapSet.new()
  end
end
