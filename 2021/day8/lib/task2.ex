defmodule Task1 do
  def solution(inputs) do
    Enum.map(inputs, &processLine/1) |> IO.inspect() |> Enum.sum()
  end

  def processLine(line) do
    String.split(line, " | ")
    |> List.last()
    |> String.split(" ")
    |> Enum.reduce(
      0,
      fn x, acc ->
        case String.length(x) do
          l when (l >= 2 and l <= 4) or l == 7 -> acc + 1
          _ -> acc
        end
      end
    )
  end
end
