defmodule Task1 do
  @spec solution({[integer()], [integer()]}) :: integer()
  def solution({left, right}) do
    Enum.zip(Enum.sort(left), Enum.sort(right))
    |> Enum.map(fn {l, r} -> abs(l - r) end)
    |> Enum.sum()
  end
end
