defmodule Task1 do
  @spec solution([]) :: integer()
  def solution(contents) do
    contents |> Enum.map(&atomize/1) |> Enum.map(&score/1) |> Enum.sum()
  end

  defp atomize({left, right}) do
    {convert(left), convert(right)}
  end

  defp convert(val) do
    cond do
      Enum.member?(["X", "A"], val) -> :rock
      Enum.member?(["Y", "B"], val) -> :paper
      Enum.member?(["Z", "C"], val) -> :scissors
    end
  end

  defp score({left, right}) do
    win = winScore(left, right)
    shape = shapeScore(right)

    win + shape
  end

  defp winScore(left, right) when left == right do
    3
  end

  defp winScore(left, right) do
    case {left, right} do
      {:rock, :paper} -> 6
      {:paper, :scissors} -> 6
      {:scissors, :rock} -> 6
      _ -> 0
    end
  end

  defp shapeScore(shape) do
    case shape do
      :rock -> 1
      :paper -> 2
      :scissors -> 3
    end
  end
end
