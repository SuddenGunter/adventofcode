defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    parse(input)
    |> Enum.map(fn x -> max_joltage(x) end)
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

  defp max_joltage(line) do
    max_first_digit =
      line
      |> Enum.take(length(line) - 1)
      |> Enum.max()

    first_digit_index = Enum.find_index(line, fn x -> x == max_first_digit end)

    max_second_digit =
      Enum.drop(line, first_digit_index + 1)
      |> Enum.take(length(line) - 1)
      |> Enum.max()

    max_first_digit * 10 + max_second_digit
  end
end
