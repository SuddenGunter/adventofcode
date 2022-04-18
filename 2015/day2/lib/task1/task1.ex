defmodule Task1 do
  @spec solution([[integer()]]) :: integer()
  def solution(dimensions) do
    dimensions |> Enum.map(fn x -> sizeOfPaper(x) end) |> Enum.sum()
  end

  @spec sizeOfPaper([integer()]) :: integer()
  def sizeOfPaper(dimensions) do
    area(dimensions) + slack(dimensions)
  end

  @spec area([integer()]) :: integer()
  def area(dimensions) do
    {l, w, h} =
      dimensions
      |> Enum.with_index()
      |> Enum.reduce({0, 0, 0}, fn {element, index}, acc ->
        put_elem(acc, index, element)
      end)

    2 * l * w + 2 * l * h + 2 * w * h
  end

  @spec slack([integer()]) :: integer()
  def slack(dimensions) do
    dimensions |> Enum.sort() |> Enum.take(2) |> Enum.reduce(1, fn x, acc -> acc * x end)
  end
end
