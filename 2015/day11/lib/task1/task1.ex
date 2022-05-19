defmodule Task1 do
  @spec solution(String.t()) :: String.t()
  def solution(line) do
    line
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

  defp includesPairs(_) do
    true
  end

  defp includesIncreasingStraight(_) do
    true
  end

  defp notIncludesBanned(line) do
    not String.contains?(line, ["i", "o", "l"])
  end
end
