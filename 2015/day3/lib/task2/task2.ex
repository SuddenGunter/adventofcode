defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(moves) do
    moves
    |> String.graphemes()
    |> Enum.reduce(
      %{
        set: MapSet.new([{0, 0}]),
        santa: {0, 0},
        robot: {0, 0},
        current: :santa
      },
      fn x, acc -> iterate(x, acc) end
    )
    |> Map.fetch(:set)
    |> elem(1)
    |> MapSet.size()
  end

  defp iterate(x, acc) do
    current = acc[:current]
    newPos = calculateNewPos(acc[current], x)

    Map.put(acc, current, newPos)
    |> Map.put(:set, MapSet.put(acc[:set], newPos))
    |> Map.put(:current, opposite(current))
  end

  defp opposite(current) do
    case current do
      :santa -> :robot
      :robot -> :santa
    end
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
