defmodule Task2 do
  @spec solution([[integer()]]) :: integer()
  def solution(reports) do
    reports
    |> Enum.map(&safe_level?/1)
    |> Enum.count(fn x -> x end)
  end

  @spec safe_level?([integer()]) :: boolean()
  def safe_level?(data) do
    res =
      Enum.with_index(data)
      |> Enum.flat_map(fn {_, i} ->
        {l, r} = Enum.split(data, i)
        [l ++ Enum.drop(r, 1)]
      end)

    (res ++ [data]) |> Enum.map(&Task1.safe_level?/1) |> Enum.any?()
  end
end
