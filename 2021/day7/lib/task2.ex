defmodule Task2 do
  def solution(inputs) do
    m = mean(inputs)

    min(
      total(inputs, round(m)),
      total(inputs, trunc(m))
    )
  end

  def total(inputs, m) do
    inputs
    |> Enum.map(fn x -> x - m end)
    |> Enum.map(fn x -> abs(x) end)
    |> Enum.map(fn x -> cost(x) end)
    |> Enum.sum()
    |> trunc
  end

  def mean(inputs) do
    Enum.sum(inputs) / length(inputs)
  end

  def cost(n) do
    (n * (n + 1) / 2) |> round
  end
end
