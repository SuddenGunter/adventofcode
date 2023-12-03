defmodule Task1 do
  @spec solution([String.t()]) :: integer()
  def solution(lines) do
    Enum.map(lines, &digits/1) |> Enum.map(&first_and_last/1) |> Enum.sum()
  end

  @spec digits(binary()) :: [integer()]
  def digits(line) do
    String.codepoints(line)
    |> Enum.map(&Integer.parse/1)
    |> Enum.filter(&match?({_, _}, &1))
    |> Enum.map(fn {x, _} -> x end)
  end

  @spec first_and_last([integer()]) :: integer()
  def first_and_last(nums) do
    Enum.at(nums, 0) * 10 + Enum.at(nums, -1)
  end
end
