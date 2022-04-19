defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(moves) do
    moves |> Enum.reduce(%{set: MapSet.new({0,0}), pos: {0,0}}, fn x, acc -> iterate(x, acc) end)
  end

  def iterate(x, acc) do
    newPos = acc[:pos]
  end
end
