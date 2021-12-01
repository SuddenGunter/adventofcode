defmodule Day2 do
  def solution(inputs, prevSumOf3, acc) do
    case length(inputs) do
      n when n <= 3 ->
        acc

      _ ->
        sumOf3 = Enum.take(inputs, 3) |> Enum.sum()

        solution(
          Enum.drop(inputs, 1),
          sumOf3,
          if prevSumOf3 < sumOf3 do
            acc + 1
          else
            acc + 0
          end
        )
    end
  end
end
