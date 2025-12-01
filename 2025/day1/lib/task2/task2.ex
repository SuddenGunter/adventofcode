defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    start = 50
    rotate(input |> String.trim() |> String.split("\n"), start, 0)
  end

  defp rotate([], _current_pos, total_zeroes) do
    total_zeroes
  end

  defp rotate([h | tail], current_pos, total_zeroes) do
    {new_pos, clicks_while_rotating} = add(h, current_pos)

    {pos, zeroes} =
      cond do
        new_pos == 100 || new_pos == 0 ->
          {0, total_zeroes + 1}

        new_pos > 99 ->
          total = if current_pos != 0, do: total_zeroes + 1, else: total_zeroes
          {new_pos - 100, total}

        new_pos < 1 ->
          total = if current_pos != 0, do: total_zeroes + 1, else: total_zeroes
          {100 + new_pos, total}

        true ->
          {new_pos, total_zeroes}
      end

    rotate(tail, pos, zeroes + clicks_while_rotating)
  end

  defp add("L" <> moves, current_pos) do
    {num, _} = Integer.parse(moves)
    real_moves = rem(num, 100)
    {current_pos - real_moves, div(num, 100)}
  end

  defp add("R" <> moves, current_pos) do
    {num, _} = Integer.parse(moves)
    real_moves = rem(num, 100)
    {current_pos + real_moves, div(num, 100)}
  end
end
