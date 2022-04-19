defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(moves) do
    moves
    |> String.graphemes()
    |> Enum.reduce(%{set: MapSet.new([{0, 0}]), pos: {0, 0}}, fn x, acc -> iterate(x, acc) end)
    |> Map.fetch(:set)
    |> elem(1)
    |> MapSet.size()
  end

  defp iterate(x, acc) do
    newPos = calculateNewPos(acc[:pos], x)

    %{
      set: MapSet.put(acc[:set], newPos),
      pos: newPos
    }
  end

  defp calculateNewPos(oldPos, modifier) do
    {
      elem(oldPos, 0) + horizontalMod(modifier),
      elem(oldPos, 1) + verticalMod(modifier)
    }
  end

  defp horizontalMod(modifier) do
    case modifier do
      "<" -> -1
      ">" -> 1
      _ -> 0
    end
  end

  defp verticalMod(modifier) do
    case modifier do
      "v" -> -1
      "^" -> 1
      _ -> 0
    end
  end
end
