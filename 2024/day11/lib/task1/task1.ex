defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    input
    |> String.split("\n", trim: true)
    |> Enum.at(0)
    |> String.split(" ", trim: true)
    |> Enum.map(&String.to_integer/1)
    |> Enum.flat_map(fn x -> blink(x, 25) end)
    |> Enum.count()
  end

  def blink(x, 0) do
    x
  end

  def blink([], _) do
    []
  end

  def blink([x | tail], num) do
    blink(x, num) ++ blink(tail, num)
  end

  def blink(x, num) when is_integer(x) do
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
