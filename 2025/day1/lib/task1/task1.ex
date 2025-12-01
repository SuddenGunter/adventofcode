defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    start = 50
    rotate(input |> String.trim() |> String.split("\n"), start, 0)
  end

  defp rotate([], _current_pos, total_zeroes) do
    total_zeroes
  end

  defp rotate([h | tail], current_pos, total_zeroes) do
    new_pos = add(h, current_pos)

    {pos, zeroes} =
      cond do
        new_pos == 100 || new_pos == 0 ->
          {0, total_zeroes + 1}

        new_pos > 99 ->
          {new_pos - 100, total_zeroes}

        new_pos < 1 ->
          {100 + new_pos, total_zeroes}

        true ->
          {new_pos, total_zeroes}
      end

    rotate(tail, pos, zeroes)
  end

  defp add("L" <> moves, current_pos) do
    {num, _} = Integer.parse(moves)
    real_moves = rem(num, 100)
    current_pos - real_moves
  end

  defp add("R" <> moves, current_pos) do
    {num, _} = Integer.parse(moves)
    real_moves = rem(num, 100)
    current_pos + real_moves
  end
end
