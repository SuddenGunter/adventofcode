defmodule Task2 do
  @spec solution([]) :: integer()
  def solution(lines) do
    lines
    |> Enum.chunk_every(3)
    |> Enum.map(&badges/1)
    |> Enum.map(&score/1)
    |> Enum.sum()
  end

  defp badges(group) do
    sets = Enum.map(group, fn x -> MapSet.new(String.graphemes(x)) end)
    first = Enum.at(sets, 0)

    Enum.drop(sets, 1)
    |> Enum.reduce(first, fn x, acc -> MapSet.intersection(x, acc) end)
    |> MapSet.to_list()
    |> Enum.at(0)
  end

  defp score(item) do
    {code, _} = String.next_codepoint(item)
    val = :binary.decode_unsigned(code)

    if item == String.upcase(item) do
      val - 38
    else
      val - 96
    end
  end
end
