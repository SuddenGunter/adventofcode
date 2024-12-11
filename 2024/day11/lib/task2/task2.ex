defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    input
    |> String.split("\n", trim: true)
    |> Enum.at(0)
    |> String.split(" ", trim: true)
    |> Enum.map(&String.to_integer/1)
    |> Enum.map(fn x -> blink(x, 75) end)
    |> Enum.sum()
  end

  def blink(x, 0) do
    length(x)
  end

  def blink([], _) do
    0
  end

  def blink([x | tail], num) do
    blink_memo(x, num) + blink(tail, num)
  end

  def blink(x, num) do
    blink_memo(x, num)
  end

  use Memoize

  defmemo blink_memo(x, num) when is_integer(x) do
    case x do
      0 ->
        blink([1], num - 1)

      x ->
        log10 = :math.log10(x) |> trunc()

        case log10 |> rem(2) do
          0 ->
            blink([x * 2024], num - 1)

          _ ->
            pow10 = :math.pow(10, (log10 / 2) |> round()) |> trunc()
            blink([div(x, pow10), rem(x, pow10)], num - 1)
        end
    end
  end
end
