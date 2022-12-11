defmodule Task2 do
  @spec solution([Monkey.t()]) :: integer()
  def solution(monkeys) do
    num = get_num(monkeys)

    Solver.solution(
      monkeys,
      10000,
      fn x -> rem(x, num) end
    )
  end

  defp get_num(monkeys) do
    Enum.map(monkeys, fn x -> x.test_number end)
    |> Enum.reduce(1, fn x, acc -> x * acc end)
  end
end
