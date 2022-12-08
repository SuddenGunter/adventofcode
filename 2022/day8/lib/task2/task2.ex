defmodule Task2 do
  @spec solution([[integer()]]) :: integer()
  def solution(input) do
    row_score = score_by_rows(input)

    transposed = transpose(input)

    col_score = score_by_rows(transposed) |> transpose()

    Enum.zip(row_score, col_score)
    |> Enum.map(fn {l, r} -> Enum.zip(l, r) |> Enum.map(fn {i, j} -> i * j end) end)
    |> List.flatten()
    |> Enum.max()
  end

  @spec score_by_rows([[integer()]]) :: [[integer()]]
  defp score_by_rows(matrix) do
    matrix
    |> Enum.map(fn x -> score_by_row([], Enum.at(x, 0), Enum.drop(x, 1)) |> List.flatten() end)
  end

  @spec score_by_row([integer()], integer(), [integer()]) :: [integer()]
  defp score_by_row(left, cur_char, right) do
    if right == [] do
      [0]
    else
      right_score = score_direction(right, cur_char)
      left_score = score_direction(left, cur_char)
      tail = score_by_row([cur_char | left], Enum.at(right, 0), Enum.drop(right, 1))
      [left_score * right_score | tail]
    end
  end

  @spec score_direction([integer()], integer()) :: integer()
  defp score_direction(trees, height) do
    if trees == [] do
      0
    else
      case Enum.find_index(trees, fn x -> x >= height end) do
        nil -> length(trees)
        n -> n + 1
      end
    end
  end

  @spec transpose([[integer()]]) :: [[integer()]]
  defp transpose(rows) do
    rows
    |> List.zip()
    |> Enum.map(&Tuple.to_list/1)
  end
end
