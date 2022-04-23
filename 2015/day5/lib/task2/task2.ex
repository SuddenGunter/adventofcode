defmodule Task2 do
  @spec solution([String.t()]) :: integer()
  def solution(lines) do
    lines |> Stream.filter(fn x -> isNice(x) end) |> Enum.count()
  end

  defp isNice(str) do
    chars = String.graphemes(str)
    containsTwoPairs(chars) and containsPairWithMiddleElement(chars)
  end

  defp containsTwoPairs(chars) do
    pairs =
      Stream.zip(chars, Stream.drop(chars, 1))
      |> Enum.map(fn {l, r} -> l <> r end)

    str = Enum.join(chars)

    pairs
    |> Enum.any?(fn x -> length(String.split(str, x)) > 2 end)
  end

  defp containsPairWithMiddleElement(chars) do
    Stream.drop(chars, 2) |> Stream.zip(chars) |> Enum.any?(fn {l, r} -> l === r end)
  end
end
