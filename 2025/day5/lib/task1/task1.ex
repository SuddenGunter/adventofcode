defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    {intervals, ingridients} =
      parse(input)
      |> IO.inspect()

    0
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
end
