defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    parse(input)
    |> solve(0)
  end

  defp parse(file) do
    String.trim(file)
    |> String.split("\n")
    |> Enum.reverse()
    |> Enum.map(fn line ->
      String.split(line)
    end)
  end

  defp solve([[] | _tail], total) do
    total
  end

  defp solve([[cmd | tail_cmds] | numbers], total) do
    {new_nubmers, operation_total} =
      case cmd do
        "*" ->
          Enum.map_reduce(numbers, 1, fn [x | xtail], acc ->
            {xtail, acc * parse_int!(x)}
          end)

        "+" ->
          Enum.map_reduce(numbers, 0, fn [x | xtail], acc ->
            {xtail, acc + parse_int!(x)}
          end)
      end

    solve([tail_cmds | new_nubmers], total + operation_total)
  end

  defp parse_int!(x) do
    {num, _} = Integer.parse(x)
    num
  end
end
