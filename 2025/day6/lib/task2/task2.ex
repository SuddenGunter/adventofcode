defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    parse(input)
    |> Enum.map(&apply_op/1)
    |> Enum.sum()
  end

  defp apply_op({"+", nums}) do
    Enum.sum(nums)
  end

  defp apply_op({"*", nums}) do
    Enum.product(nums)
  end

  defp parse(file) do
    [ops_row | numbers] =
      file
      |> String.split("\n")
      |> Enum.map(&String.graphemes/1)
      |> Enum.reverse()

    #  <> " " is a hack to simplify parsing of last column
    ops = parse_ops(ops_row ++ [" "], [])
    parse_nums(ops, numbers |> Enum.reverse(), [])
  end

  defp parse_nums([], _numbers, parsed) do
    parsed
  end

  defp parse_nums([{op, len} | tail_ops], numbers, parsed) do
    {next_numbers, op_numbers} =
      Enum.map_reduce(numbers, List.duplicate(0, len), fn x, acc ->
        op_row =
          Enum.take(x, len)
          |> Enum.map(fn
            " " -> :nan
            num -> parse_int!(num)
          end)

        next_acc =
          Enum.zip_with(
            [acc, op_row],
            fn
              [l, :nan] -> l
              [l, r] -> l * 10 + r
            end
          )

        {Enum.drop(x, len + 1), next_acc}
      end)

    parse_nums(tail_ops, next_numbers, [{op, op_numbers} | parsed])
  end

  defp parse_ops([], ops) do
    ops |> Enum.reverse()
  end

  defp parse_ops([op | row], ops) do
    len =
      row
      |> Enum.take_while(fn x -> x == " " end)
      |> length()

    parse_ops(row |> Enum.drop(len), [{op, len} | ops])
  end

  defp parse_int!(x) do
    {num, _} = Integer.parse(x)
    num
  end
end
