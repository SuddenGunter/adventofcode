defmodule Task2 do
  @spec solution({[integer()], [integer()]}) :: integer()
  def solution({left, right}) do
    num_occurrences =
      Enum.reduce(right, %{}, fn x, acc ->
        Map.update(acc, x, 1, fn c -> c + 1 end)
      end)

    Enum.map(left, fn x ->
      Map.get(num_occurrences, x, 0) * x
    end)
    |> Enum.sum()
  end
end
