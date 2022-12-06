defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(contents) do
    contents |> String.graphemes() |> firstUniquePos(4)
  end

  defp firstUniquePos(letters, pos) do
    firstFour = Enum.take(letters, 4) |> Enum.uniq()

    if length(firstFour) == 4 do
      pos
    else
      firstUniquePos(Enum.drop(letters, 1), pos + 1)
    end
  end
end
