defmodule Task1 do
  @spec solution([String.t()]) :: integer()
  def solution(lines) do
    lines |> Stream.filter(fn x -> isNice(x) end) |> Enum.count()
  end

  defp isNice(str) do
    chars = String.graphemes(str)
    containsVowels(chars) and containsDuplicate(chars) and containsNoForbiddenContent(chars)
  end

  defp containsVowels(chars) do
    vowels = ["a", "e", "i", "o", "u"]
    chars |> Stream.filter(fn x -> Enum.any?(vowels, fn m -> m == x end) end) |> Enum.count() >= 3
  end

  defp containsDuplicate(chars) do
    Stream.drop(chars, 1) |> Stream.zip(chars) |> Enum.any?(fn {l, r} -> l === r end)
  end

  defp containsNoForbiddenContent(chars) do
    forbidden = ["ab", "cd", "pq", "xy"]

    contains =
      Enum.zip(chars, Enum.drop(chars, 1))
      |> Enum.any?(fn {l, r} -> Enum.any?(forbidden, fn x -> x == l <> r end) end)

    not contains
  end
end
