defmodule Task2 do
  @spec solution([[integer()]]) :: integer()
  def solution(dimensions) do
    dimensions |> Enum.map(fn x -> sizeOfRibbon(x) end) |> Enum.sum()
  end

  @spec sizeOfRibbon([integer()]) :: integer()
  def sizeOfRibbon(dimensions) do
    wrap(dimensions) + bow(dimensions)
  end

  @spec wrap([integer()]) :: integer()
  def wrap(dimensions) do
    dimensions |> Enum.reduce(1, fn x, acc -> acc * x end)
  end

  @spec bow([integer()]) :: integer()
  def bow(dimensions) do
    half = dimensions |> Enum.sort() |> Enum.take(2) |> Enum.reduce(0, fn x, acc -> acc + x end)
    2 * half
  end
end
