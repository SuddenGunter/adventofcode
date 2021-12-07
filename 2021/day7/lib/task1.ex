defmodule Task1 do
  def solution(inputs) do
    m = median(inputs)

    inputs |> Enum.map(fn x -> x - m end) |> Enum.map(fn x -> abs(x) end) |> Enum.sum() |> trunc
  end

  def median(list) do
    len = length(list)
    sorted = Enum.sort(list)
    mid = div(len, 2)

    if rem(len, 2) == 0,
      do: (Enum.at(sorted, mid - 1) + Enum.at(sorted, mid)) / 2,
      else: Enum.at(sorted, mid)
  end
end
