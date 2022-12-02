defmodule Task2 do
  @spec solution([]) :: integer()
  def solution(contents) do
    contents |> Enum.map(&takeAction/1) |> Task1.solution()
  end

  defp takeAction({left, right}) do
    action =
      case right do
        "X" -> lose(left)
        "Y" -> draw(left)
        "Z" -> win(left)
      end

    {left, action}
  end

  defp win(shape) do
    case shape do
      "A" -> "Y"
      "B" -> "Z"
      "C" -> "X"
    end
  end

  defp lose(shape) do
    case shape do
      "A" -> "Z"
      "B" -> "X"
      "C" -> "Y"
    end
  end

  defp draw(shape) do
    case shape do
      "A" -> "X"
      "B" -> "Y"
      "C" -> "Z"
    end
  end
end
