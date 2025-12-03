defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    parse(input)
    |> Enum.map(fn x -> max_joltage(x, 12, 0) end)
    |> Enum.sum()
  end

  defp parse(input) do
    String.trim(input)
    |> String.split("\n")
    |> Enum.map(fn line ->
      String.graphemes(line)
      |> Enum.map(fn x ->
        {num, _} = Integer.parse(x)
        num
      end)
    end)
  end

  defp max_joltage(line, need_more_digits, joltage) when length(line) == need_more_digits do
    Enum.reduce(line, joltage, fn x, acc -> acc * 10 + x end)
  end

  defp max_joltage(_line, 0, joltage) do
    joltage
  end

  defp max_joltage(line, need_more_digits, joltage) do
    max_digit =
      line
      |> Enum.take(length(line) - need_more_digits + 1)
      |> Enum.max()

    max_digit_index = Enum.find_index(line, fn x -> x == max_digit end)

    max_joltage(
      Enum.drop(line, max_digit_index + 1),
      need_more_digits - 1,
      joltage * 10 + max_digit
    )
  end
end
