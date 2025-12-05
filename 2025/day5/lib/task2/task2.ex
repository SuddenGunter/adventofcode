defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    {intervals, _ingridients} =
      parse(input)

    merge_overlaps(intervals)
    |> Enum.reduce(0, fn {l, r}, acc ->
      acc + r - l + 1
    end)
  end

  defp parse(input) do
    String.trim(input)
    |> String.split("\n")
    |> Enum.reduce({[], []}, fn x, {intervals, ingridients} ->
      cond do
        String.contains?(x, "-") ->
          {l, r} =
            String.split(x, "-")
            |> Enum.map(fn n ->
              {num, _} = Integer.parse(n)
              num
            end)
            |> List.to_tuple()

          {[{l, r} | intervals], ingridients}

        x == "" ->
          {intervals, ingridients}

        true ->
          {num, _} = Integer.parse(x)
          {intervals, [num | ingridients]}
      end
    end)
    |> Tuple.to_list()
    |> Enum.map(fn x -> Enum.reverse(x) end)
    |> List.to_tuple()
  end

  defp merge_overlaps(intervals) do
    Enum.sort(intervals)
    |> Enum.reduce([], fn
      {l, r}, [{last_l, last_r} | prev] ->
        cond do
          l >= last_l and l <= last_r ->
            [{min(l, last_l), max(r, last_r)} | prev]

          true ->
            [{l, r}, {last_l, last_r}] ++ prev
        end

      {l, r}, [] ->
        [{l, r}]
    end)
  end
end
