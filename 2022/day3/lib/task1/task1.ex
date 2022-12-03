defmodule Task1 do
  @spec solution([]) :: integer()
  def solution(lines) do
    lines
    |> Enum.map(&split/1)
    |> Enum.map(&duplicates/1)
    |> Enum.map(&score/1)
    |> Enum.sum()
  end

  defp split(rucksack) do
    midpoint = div(String.length(rucksack), 2)
    {String.slice(rucksack, 0, midpoint), String.slice(rucksack, midpoint, midpoint)}
  end

  defp duplicates({left, right}) do
    set = MapSet.new(String.graphemes(right))
    findDuplicateRec(String.graphemes(left), set)
  end

  defp findDuplicateRec(left, set) do
    [head | tail] = left

    if MapSet.member?(set, head) do
      head
    else
      findDuplicateRec(tail, set)
    end
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
