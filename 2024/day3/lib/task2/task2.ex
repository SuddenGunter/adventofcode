defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(memory) do
    Regex.scan(~r/mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)/, memory)
    |> interpret(:enabled, 0)
  end

  defp interpret([["do()"] | t], _, acc) do
    interpret(t, :enabled, acc)
  end

  defp interpret([["don't()"] | t], _, acc) do
    interpret(t, :disabled, acc)
  end

  defp interpret([], _, acc) do
    acc
  end

  defp interpret([[_, _, _] | t], :disabled, acc) do
    interpret(t, :disabled, acc)
  end

  defp interpret([[_, a, b] | t], :enabled, acc) do
    interpret(t, :enabled, acc + String.to_integer(a) * String.to_integer(b))
  end
end
