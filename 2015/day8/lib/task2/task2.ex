defmodule Task2 do
  @spec solution([String.t()]) :: integer()
  def solution(lines) do
    lines
    |> Enum.reduce(0, fn x, acc ->
      acc + encodedLength(x) - String.length(x)
    end)
  end

  defp encodedLength(str) do
    str
    |> String.replace("\\", "\\\\")
    |> String.replace("\"", "\\\"")
    |> quoteStr
    |> valueLength
  end

  defp quoteStr(str) do
    "\"#{str}\""
  end

  defp valueLength(str) do
    String.graphemes(str) |> Enum.count()
  end
end
