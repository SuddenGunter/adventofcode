defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(memory) do
    Regex.scan(~r/mul\((\d{1,3}),(\d{1,3})\)/, memory)
    |> Enum.map(fn [_, a, b] -> String.to_integer(a) * String.to_integer(b) end)
    |> Enum.sum()
  end
end
