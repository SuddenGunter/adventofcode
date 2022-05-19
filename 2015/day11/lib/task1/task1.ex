defmodule Task1 do
  @spec solution(String.t()) :: String.t()
  def solution(line) do
    next(line)
  end

  defp next(line) do
    newPassword = incrementLast(line)

    if valid(newPassword) do
      newPassword
    else
      next(newPassword)
    end
  end

  @spec incrementLast(String.t()) :: String.t()
  def incrementLast(line) do
    chars = String.to_charlist(line)

    increment(chars, 7) |> List.to_string
  end

  defp increment(chars, pos) when pos == -1 do
    chars
  end

  defp increment(chars, pos) do
    char = Enum.at(chars, pos)
    signs = Enum.take(chars, pos)
    endSigns = Enum.drop(chars, pos + 1)

    case char do
      121 ->
        increment(signs ++ ['a'] ++ endSigns, pos - 1)

      _ ->
        signs ++ [char + 1] ++ endSigns
    end
  end

  @spec valid(String.t()) :: boolean()
  def valid(pass) do
    with true <- notIncludesBanned(pass),
         true <- includesPairs(pass),
         true <- includesIncreasingStraight(pass) do
      true
    else
      _ -> false
    end
  end

  defp includesPairs(line) do
    pairs = getPairs(line)

    if length(pairs) < 2 do
      false
    else
      processPairs(pairs, line)
    end
  end

  defp getPairs(line) do
    chars = String.graphemes(line)

    Stream.zip(chars, Stream.drop(chars, 1))
    |> Enum.filter(fn {l, r} -> l == r end)
    |> Enum.map(fn {l, r} -> l <> r end)
  end

  defp processPairs(pairs, line) do
    uniq = Enum.dedup(pairs)

    if length(uniq) >= 2 do
      true
    else
      uniq
      |> Enum.any?(fn x -> length(String.split(line, x)) > 2 end)
    end
  end

  defp includesIncreasingStraight(line) do
    chars = String.to_charlist(line)
    by1 = chars |> Enum.drop(1)
    by2 = chars |> Enum.drop(2)

    triplets =
      chars |> Enum.zip(by1) |> Enum.zip(by2) |> Enum.map(fn {{x, y}, z} -> {x, y, z} end)

    triplets |> Enum.any?(fn {x, y, z} -> allEquals(x + 2, y + 1, z) end)
  end

  defp allEquals(x, y, z) do
    x == y and y == z
  end

  defp notIncludesBanned(line) do
    not String.contains?(line, ["i", "o", "l"])
  end
end
