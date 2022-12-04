defmodule Task1 do
  @spec solution([]) :: integer()
  def solution(contents) do
    contents
    |> Enum.map(&sortIntervals/1)
    |> Enum.map(&is_included/1)
    |> Enum.sum()
  end

  defp sortIntervals({l, r}) do
    l_len = elem(l, 1) - elem(l, 0)
    r_len = elem(r, 1) - elem(r, 0)
    if l_len >= r_len, do: {l, r}, else: {r, l}
  end

  defp is_included({l, r}) do
    {l_start, l_end} = {elem(l, 0), elem(l, 1)}
    {r_start, r_end} = {elem(r, 0), elem(r, 1)}
    if l_start <= r_start and l_end >= r_end, do: 1, else: 0
  end
end
