defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(line) do
    iterate(line, 0) |> String.length()
  end

  defp iterate(line, depth) when depth == 40 do
    line
  end

  defp iterate(line, depth) do
    iterate(next(line), depth + 1)
  end

  @spec next(String.t()) :: integer()
  def next(line) do
    line
    |> String.graphemes()
    |> Enum.reduce(
      %{result: "", currentChar: "", currentCount: 0},
      fn x, acc ->
        if acc[:currentChar] != x do
          %{
            result: nextResult(acc[:result], acc[:currentChar], acc[:currentCount]),
            currentChar: x,
            currentCount: 1
          }
        else
          %{
            result: acc[:result],
            currentChar: x,
            currentCount: acc[:currentCount] + 1
          }
        end
      end
    )
    |> nextResult()
  end

  defp nextResult(acc) do
    nextResult(acc[:result], acc[:currentChar], acc[:currentCount])
  end

  defp nextResult(prevResult, _, count) when count == 0 do
    prevResult
  end

  defp nextResult(prevResult, char, count) do
    prevResult <> Integer.to_string(count) <> char
  end
end
